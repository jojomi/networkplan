package main

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

// EnvPlan encapsulates the environment for the CLI plan handler.
type EnvPlan struct {
	ConfigFilename string
	OutputFilename string

	PrintAllIPv4 bool
}

// ParseFrom reads the state from a given cobra command and its args.
func (e *EnvPlan) ParseFrom(command *cobra.Command) error {
	var (
		f   = command.Flags()
		err error
	)

	e.ConfigFilename, err = f.GetString("config")
	if err != nil {
		return err
	}
	e.ConfigFilename, err = WithExpandedHome(e.ConfigFilename)
	if err != nil {
		return err
	}

	e.PrintAllIPv4, err = f.GetBool("print-all-ipv4")
	if err != nil {
		return err
	}

	return nil
}

func WithExpandedHome(input string) (string, error) {
	if !strings.HasPrefix(input, "~/") {
		return input, nil
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		return input, err
	}
	return filepath.Join(dirname, input[2:]), nil
}
