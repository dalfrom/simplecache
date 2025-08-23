package tcp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	//"os"

	//"github.com/dalfrom/tempodb/pkg/logger"
	"github.com/dalfrom/tempodb/pkg/cache"
	"github.com/dalfrom/tempodb/pkg/telemetry"
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
		fmt.Println("Error starting TempoDB server:", err)
		return fmt.Errorf("failed to start TempoDB server: %w", err)
	}

	cache.CreateCache()

	fmt.Printf("TempoDB server listening on :%d\n", t.Port)
	ctx := context.Background()
	for {
		conn, _ := t.Net.Accept()
		go handleConn(ctx, conn)
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
