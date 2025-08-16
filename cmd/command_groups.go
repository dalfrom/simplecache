package cmd

import (
	"github.com/dalfrom/tempodb/cmd/config"
	"github.com/dalfrom/tempodb/cmd/connect"
	"github.com/dalfrom/tempodb/cmd/start"
	"github.com/dalfrom/tempodb/cmd/stop"

	"github.com/spf13/cobra"
)

type CommandGroup struct {
	Message  string
	Commands []*cobra.Command
}

type CommandGroups []CommandGroup

var commandGroups = CommandGroups{
	{
		Message: "Basic Commands",
		Commands: []*cobra.Command{
			start.StartCmd,
			stop.StopCmd,
			start.StatusCmd,
			VersionCmd,
		},
	},
	{
		Message: "Configuration Commands",
		Commands: []*cobra.Command{
			config.ResetConfigCmd,
			config.SetConfigCmd,
			config.GetConfigCmd,
		},
	},
}

func addCommandGroups() {
	// We add at the very start the ability to connect via REPL Loop
	RootCmd.AddCommand(connect.ConnectCmd)

	// Then, all secondary commands
	for _, group := range commandGroups {
		RootCmd.AddCommand(group.Commands...)
	}
}
