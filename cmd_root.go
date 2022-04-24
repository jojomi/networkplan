package main

import "github.com/spf13/cobra"

func getCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "networkplan",
		Short: "networkplan",
	}

	rootCmd.AddCommand(
		getCmdPlan(),
		getCmdHostsfile(),
	)

	f := rootCmd.PersistentFlags()
	f.StringP("config", "c", "~/.networkplan/network.yml", "Network config")

	return rootCmd
}
