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
	env := EnvHostsfile{}
	err := env.ParseFrom(cmd, args)
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse params")
	}
	handleHostsfile(env)
}

func handleHostsfile(env EnvHostsfile) {
	templateData, err := ioutil.ReadFile("templates/hosts")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	config, err := LoadNetworkConfigFromFile(env.ConfigFilename)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	funcs := template.FuncMap{
		"join": strings.Join,
		"newline": func(count ...int) string {
			var c int
			if len(count) == 0 {
				c = 1
			} else {
				c = count[0]
			}
			return strings.Repeat("\n", c)
		},
	}
	output, err := strtpl.EvalWithFuncMap(string(templateData), funcs, config)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	fmt.Println(strings.TrimSpace(output))
}
