package tcp

import (
	"fmt"
	"log/slog"
	"net"

	//"github.com/dalfrom/simplecache/pkg/logger"
	"github.com/dalfrom/simplecache/pkg/cache"
	"github.com/dalfrom/simplecache/pkg/telemetry"
)

type ServerCache struct {
	// The main listener for incoming TCP connections
	Net net.Listener

	// The list of currently opened connections,
	// which will be loop-closed on shutdown
	OpenConnections []net.Conn

	// Port
	Port int

	// The configuration
	Cfgs map[string]any

	// Logging
	Logger *slog.Logger

	// This manages telemetry
	Telemetry telemetry.Telemetry
}

func (t ServerCache) Start() (err error) {
	t.Net, err = net.Listen("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		fmt.Println("Error starting simplecache server:", err)
		return fmt.Errorf("failed to start simplecache server: %w", err)
	}

	cache.CreateCache()

	fmt.Printf("simplecache server listening on :%d\n", t.Port)
	for {
		conn, _ := t.Net.Accept()
		go handleConn(&cache.SimpleCache, conn)
	}
}

func (t ServerCache) Stop(force bool) {
	if !force {
		for _, conn := range t.OpenConnections {
			conn.Close()
		}
	}

	t.Net.Close()
}
