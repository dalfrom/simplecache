package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Example: "tempodb version",
	Short:   "Print the version of TempoDB",
	Run:     getVersion,
}

func getVersion(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile("version.json")
	if err != nil {
		fmt.Println("Error reading version file:", err)
		return
	}

	var versionData map[string]string
	if err := json.Unmarshal(data, &versionData); err != nil {
		fmt.Println("Error unmarshalling version data:", err)
		return
	}

	fmt.Println("TempoDB version:", versionData["version"])
}
