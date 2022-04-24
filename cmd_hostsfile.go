package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/jojomi/strtpl"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func getCmdHostsfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hostsfile",
		Short: "Export the network hosts in hostsfile format",
		Run:   cmdHostsfileHandler,
	}
	return cmd
}

func cmdHostsfileHandler(cmd *cobra.Command, args []string) {
	templateData, err := ioutil.ReadFile("templates/hosts")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	config, err := LoadNetworkConfigFromFile("network.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	funcs := template.FuncMap{
		"join": strings.Join,
	}
	output, err := strtpl.EvalWithFuncMap(string(templateData), funcs, config)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	fmt.Println(output)
}
