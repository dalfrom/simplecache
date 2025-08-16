package connect

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmdServer "github.com/dalfrom/tempodb/pkg/cmd"
	"github.com/spf13/cobra"
)

var (
	user string
	pass string
	host string
	port int
)

// ConnectCmd represents the connect command
var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the TempoDB server",
	Run: func(cmd *cobra.Command, args []string) {
		// cmdServer.Connect(user, pass, host, port)
		handlePasswordAuth()
	},
}

func init() {
	// Add flags to the connect command
	ConnectCmd.Flags().StringVarP(&user, "user", "u", "", "Username")
	ConnectCmd.Flags().StringVarP(&pass, "password", "p", "", "Password")
	if err := ConnectCmd.MarkFlagRequired("user"); err != nil {
		panic(err)
	}
	ConnectCmd.Flags().StringVar(&host, "host", "127.0.0.1", "Server host [defaults to 127.0.0.1]")
	ConnectCmd.Flags().IntVar(&port, "port", 4000, "Server port [defaults to 4000]")
}

func handlePasswordAuth() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("password> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "quit" || line == "exit" {
			fmt.Println("Aborted connection")
			break
		}

		authenticated, err := cmdServer.CheckConnection(user, line, host, port)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if authenticated {
			fmt.Println("Authentication successful!")
			pass = line
			break
		} else {
			fmt.Println("Authentication failed.")
		}
	}

	cmdServer.Connect(user, pass, host, port)
}
