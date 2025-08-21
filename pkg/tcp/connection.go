package tcp

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/dalfrom/tempodb/pkg/scl"
	"github.com/dalfrom/tempodb/pkg/tcp/security"
)

func handleConn(_ context.Context, conn net.Conn) {
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

		err = scl.Extract(line)
		if err != nil {
			fmt.Println("Error extracting SCL:", err)
			io.WriteString(conn, err.Error()+"\n")
			return
		}
	}
}
