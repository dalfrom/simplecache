/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const longDescription = `Tempo controls the TempoDB Database system.`

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tempodb",
	Short: "A brief description of your application",
	Long:  longDescription,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	addCommandGroups()
}
