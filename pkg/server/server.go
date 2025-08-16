package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/dalfrom/tempodb/pkg/server/security"
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
	if err := security.Authenticate(conn); err != nil {
		fmt.Println("Authentication error:", err)
		io.WriteString(conn, err.Error()+"\n")
		return
	}

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)

		conn.Write([]byte("Line that was received: " + line))
	}
}
