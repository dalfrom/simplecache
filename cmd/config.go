package cmd

import (
	"github.com/spf13/cobra"
)

// SetConfigCmd represents the set config command
var SetConfigCmd = &cobra.Command{
	Use:   "set-config <key>=<value> [<key>=<value> ...]",
	Short: "Set one or more configuration values",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("Invalid arguments. Usage: set-config <key>=<value> [<key>=<value> ...]")
			return
		}
		for _, arg := range args {
			kv := splitKeyValue(arg)
			if kv == nil {
				println("Invalid argument:", arg)
				continue
			}
			key, value := kv[0], kv[1]

			println("Setting TempoDB configuration:", key, "=", value)
		}
	},
}

// GetConfigCmd represents the get config command
var GetConfigCmd = &cobra.Command{
	Use:   "get-config <key>",
	Short: "Get the value of a specific configuration key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			println("Invalid arguments. Usage: get-config <key>")
			return
		}
		key := args[0]
		println("Getting TempoDB configuration for key:", key)
	},
}

// ResetConfigCmd represents the reset config command
var ResetConfigCmd = &cobra.Command{
	Use:   "reset-config",
	Short: "Reset the TempoDB configuration to default values",
	Run: func(cmd *cobra.Command, args []string) {
		println("Resetting TempoDB configuration to default values")
	},
}

// splitKeyValue splits a string of the form key=value into a [2]string slice.
func splitKeyValue(arg string) []string {
	kv := make([]string, 2)
	for i, c := range arg {
		if c == '=' {
			kv[0] = arg[:i]
			kv[1] = arg[i+1:]
			return kv
		}
	}
	return nil // return nil if no '=' found
}
