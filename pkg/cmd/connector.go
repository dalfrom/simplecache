package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func CheckConnection(user, password, host string, port int) (bool, error) {
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	// Send authentication request
	_, _ = fmt.Fprintf(conn, "AUTH %s %s\n", user, password)

	// Read response
	resp := make([]byte, 4096)
	n, _ := conn.Read(resp)
	fmt.Println(string(resp[:n]))

	// TODO:
	// This is currently blocked due to missing the mechanism to load the data
	// from the database tables upon initialization

	return true, nil
}

func Connect(user, password, host string, port int) {
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	fmt.Printf("Connected to TempoDB at %s as %s\n", addr, user)

	// Start REPL loop
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to tempodb SQL interface")
	for {
		fmt.Print("tempodb> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "quit" || line == "exit" {
			fmt.Println("Bye!")
			break
		}

		if line == "clear" {
			fmt.Print("\033[H\033[2J")
			continue
		}

		// send command to server
		_, _ = conn.Write([]byte(line + "\n"))

		// read response
		resp := make([]byte, 4096)
		n, _ := conn.Read(resp)
		fmt.Println(string(resp[:n]))
	}
}
