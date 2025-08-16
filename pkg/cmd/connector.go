package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func plain(conn net.Conn, r *bufio.Reader, user, pass string) (authenticated bool) {
	b64 := base64.StdEncoding.EncodeToString([]byte("\x00" + user + "\x00" + pass))
	fmt.Fprintln(conn, b64)

	resp, _ := r.ReadString('\n')
	if strings.HasPrefix(resp, "OK") {
		authenticated = true
	} else {
		fmt.Println("Authentication failed")
	}

	return
}

func Connect(user, password, host string, port int) {
	addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	fmt.Println("Server:", strings.TrimSpace(line))

	if !strings.HasPrefix(line, "SASL:") {
		fmt.Println("Unexpected server response:", line)
		return
	}

	fmt.Fprintf(conn, "AUTH %s\n", "PLAIN")
	if !plain(conn, r, user, password) {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to tempodb SQL interface")

	// Start REPL loop
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
		fmt.Fprintf(conn, "%s\n", line)

		// read response
		resp := make([]byte, 4096)
		n, _ := conn.Read(resp)
		fmt.Println(string(resp[:n]))
	}
}
