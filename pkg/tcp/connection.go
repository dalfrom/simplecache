package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/dalfrom/simplecache/pkg/cache"
	"github.com/dalfrom/simplecache/pkg/scl"
	"github.com/dalfrom/simplecache/pkg/tcp/security"
)

func handleConn(conn net.Conn, config *ServerCache, cache *cache.Cache) {
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

		// Before writing to the cache, we write to the wal
		config.WAL.WriteToWal(line)

		// Handling the operation that was parsed by SCL
		handleOperation(conn, line, cache)
	}
}

func handleOperation(conn net.Conn, line string, cache *cache.Cache) {
	// Extracting the data via SCL
	err := scl.Extract(line)
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
	default:
		io.WriteString(conn, "Unknown operation\n")
	}
}
