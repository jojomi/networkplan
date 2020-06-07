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

	return rootCmd
}
