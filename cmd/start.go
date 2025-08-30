package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dalfrom/simplecache/pkg/tcp"
)

const defaultPort int = 4000
const defaultConfigPath string = "/etc/simplecache/simplecache.toml"

var (
	port              int
	configurationPath string
)

func main() {
	// There are no provided config and port. Using the default ones
	if len(os.Args) < 2 {
		configurationPath = defaultConfigPath
		port = defaultPort
	} else if len(os.Args) == 2 {
		// There is only the config provided. Using the default port
		configurationPath = os.Args[1]
		port = defaultPort
	} else if len(os.Args) == 3 {
		// Both config and port are provided
		configurationPath = os.Args[1]

		iPort, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid port number. Using default port.")
			port = defaultPort
		} else {
			port = iPort
		}
	}

	serverCache := &tcp.ServerCache{Port: port}
	serverCache.Start(configurationPath)

	// Handle graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down gracefully…")
		serverCache.Stop(false)
		os.Exit(0)
	}()
}
