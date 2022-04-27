package main

import (
	"errors"
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
		"space": func(count ...int) string {
			var c int
			if len(count) == 0 {
				c = 1
			} else {
				c = count[0]
			}
			return strings.Repeat(" ", c)
		},
		"add":  templateAdd,
		"dict": templateParamDict,
	}
	output, err := strtpl.EvalWithFuncMap(string(templateData), funcs, config)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	fmt.Println(strings.TrimSpace(output))
}

func templateParamDict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func templateAdd(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
