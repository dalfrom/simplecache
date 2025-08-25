package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/dalfrom/tempodb/pkg/cache"
	"github.com/dalfrom/tempodb/pkg/scl"
	"github.com/dalfrom/tempodb/pkg/tcp/security"
)

func handleConn(cache *cache.Cache, conn net.Conn) {
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

		fmt.Fprintf(conn, "Line that was received: %s\n", line)

		err = scl.Extract(line)
		if err != nil {
			fmt.Println("Error extracting SCL:", err)
			fmt.Fprintf(conn, err.Error()+"\n")
			return
		}

		switch scl.Type {
		case "SET":
			cache.Set(scl.Collection, scl.Key, scl.Value)
		case "GET":
			value, found := cache.Get(scl.Collection, scl.Key)
			if !found {
				io.WriteString(conn, "Key not found\n")
				return
			}
			fmt.Fprintf(conn, "Value: %v\n", value)
		case "DELETE":
			cache.Delete(scl.Collection, scl.Key)
		case "TRUNCATE":
			cache.Truncate(scl.Collection)
		case "DROP":
			cache.Drop(scl.Collection)
		case "UPDATE":
			_, found := cache.Get(scl.Collection, scl.Key)
			if !found {
				cache.Set(scl.Collection, scl.Key, scl.Value)
				return
			}

			cache.Update(scl.Collection, scl.Key, scl.Value)
		}
	}
}
