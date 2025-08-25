/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dalfrom/simplecache/pkg/tcp"
)

var (
	ServerCache *tcp.ServerCache
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "simplecache",
	Short: "A no-brainer caching solution",
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
