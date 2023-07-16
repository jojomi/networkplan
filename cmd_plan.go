package main

import (
	_ "embed"
	"fmt"
	"github.com/jojomi/strtpl"
	htmlTemplate "html/template"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	flagOptionsPrintAllIPv4s bool
)

func getCmdPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Generate a network plan in an HTML file",
		Run:   cmdPlanHandler,
	}

	f := cmd.Flags()
	f.Bool("print-all-ipv4", false, "Also print unused IPv4 addresses")

	return cmd
}

type PlanExportOptions struct {
	PrintAllIPv4s bool
}

func cmdPlanHandler(cmd *cobra.Command, args []string) {
	env := EnvPlan{}
	err := env.ParseFrom(cmd)
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse params")
	}
	handlePlan(env)
}

//go:embed "templates/plan.html"
var templatePlanData []byte

func handlePlan(env EnvPlan) {
	config, err := LoadNetworkConfigFromFile(env.ConfigFilename)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	exportOptions := PlanExportOptions{
		PrintAllIPv4s: flagOptionsPrintAllIPv4s,
	}
	funcMap := htmlTemplate.FuncMap{
		"join": strings.Join,
		"add":  templateAdd,
		"dict": templateParamDict,
	}
	renderedOutput, err := strtpl.EvalHTMLWithFuncMap(string(templatePlanData), funcMap, struct {
		Config        *NetworkConfig
		ExportOptions PlanExportOptions
	}{
		Config:        config,
		ExportOptions: exportOptions,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("rendering plan template failed")
	}

	fmt.Println(renderedOutput)
}
