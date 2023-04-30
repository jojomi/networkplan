package main

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

// EnvHostsfile encapsulates the environment for the CLI hostsfile handler.
type EnvHostsfile struct {
	ConfigFilename string
}

// ParseFrom reads the state from a given cobra command and its args.
func (e *EnvHostsfile) ParseFrom(command *cobra.Command) error {
	var (
		f   = command.Flags()
		err error
	)

	e.ConfigFilename, err = f.GetString("config")
	if err != nil {
		return err
	}

	// expand home dir
	if strings.HasPrefix(e.ConfigFilename, "~/") {
		dirname, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		e.ConfigFilename = filepath.Join(dirname, e.ConfigFilename[2:])
	}

	return nil
}
