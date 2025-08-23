package cmd

import (
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
			StartCmd,
			StopCmd,
			StatusCmd,
			VersionCmd,
		},
	},
	{
		Message: "Configuration Commands",
		Commands: []*cobra.Command{
			ResetConfigCmd,
			SetConfigCmd,
			GetConfigCmd,
		},
	},
}

func addCommandGroups() {
	// We add at the very start the ability to connect via REPL Loop
	RootCmd.AddCommand(ConnectCmd)

	// Then, all secondary commands
	for _, group := range commandGroups {
		RootCmd.AddCommand(group.Commands...)
	}
}
