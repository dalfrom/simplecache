package cmd

import (
	"github.com/spf13/cobra"
)

// StopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:   "stop [flags]",
	Short: "Stop the TempoDB server",
	Long: `This command stops the TempoDB server gracefully, ensuring all operations are completed before shutdown.

	Any ongoing operations will be allowed to finish before the server stops. If you want to stop the server immediately, you can use the --force flag.`,
	Example: `tempodb stop --force`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")

		ServerCache.Stop(force)
	},
}

func init() {
	// Add a flag to force stop the server
	StopCmd.Flags().BoolP("force", "f", false, "Force stop the TempoDB server without waiting for ongoing operations to complete")
}
