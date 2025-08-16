package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Serve(port int) {
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	fmt.Printf("TempoDB server running on :%d\n", port)

	for {
		conn, _ := ln.Accept()
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)

		if strings.Contains(line, "AUTH") {
			// Handle authentication
			parts := strings.SplitN(line, " ", 3)
			if len(parts) != 3 {
				conn.Write([]byte("ERROR: Invalid AUTH command format\n"))
				return
			}
			user, password := parts[1], parts[2]
			fmt.Printf("Received authentication request for user-pass: %s-%s\n", user, password)
			conn.Write([]byte("OK: Authentication successful\n"))

			// We break because we will reconnect with the new credentials later, as this is a "double connection"
			break
		}

		switch line {
		case "status":
			conn.Write([]byte("OK: TempoDB is running"))
		case "ping":
			conn.Write([]byte("PONG"))
		default:
			conn.Write([]byte("Unknown command: " + line))
		}
	}
}
