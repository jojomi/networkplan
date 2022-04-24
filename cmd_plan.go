package main

import (
	htmlTemplate "html/template"
	"path/filepath"
	"strings"

	"github.com/hexops/valast"
	"github.com/jojomi/tplrender"
	"github.com/pkg/browser"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	flagPlanOpen             bool
	flagOptionsPrintAllIPv4s bool
)

func getCmdPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Generate a network plan in an HTML file",
		Run:   cmdPlanHandler,
	}

	f := cmd.PersistentFlags()
	f.BoolVarP(&flagPlanOpen, "open", "o", false, "output generated document")
	f.BoolVarP(&flagOptionsPrintAllIPv4s, "print-all-ipv4s", "", false, "Also print unused IPv4 addresses")

	return cmd
}

type PlanExportOptions struct {
	PrintAllIPv4s bool
}

func cmdPlanHandler(cmd *cobra.Command, args []string) {
	config, err := LoadNetworkConfigFromFile("network.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Trace().Msg(valast.String(config))

	templateFile := "plan.html"
	opts := tplrender.Options{
		TemplateDir:      "templates",
		TemplateFilename: templateFile,
		OutputDir:        "build",
		OutputFilename:   templateFile,
	}
	exportOptions := PlanExportOptions{
		PrintAllIPv4s: flagOptionsPrintAllIPv4s,
	}
	funcMap := htmlTemplate.FuncMap{
		"join": strings.Join,
	}
	err = tplrender.HTMLTemplateWithFuncMap(opts, funcMap, struct {
		Config        *NetworkConfig
		ExportOptions PlanExportOptions
	}{
		Config:        config,
		ExportOptions: exportOptions,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	if flagPlanOpen {
		browser.OpenURL(filepath.Join(opts.OutputDir, opts.OutputFilename))
	}
}
