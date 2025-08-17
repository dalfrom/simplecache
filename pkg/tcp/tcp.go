package tcp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	//"os"

	//"github.com/dalfrom/tempodb/pkg/logger"
	"github.com/dalfrom/tempodb/pkg/telemetry"
	"github.com/dalfrom/tempodb/pkg/wal"
)

type Tempo struct {
	// The main listener for incoming TCP connections
	Net net.Listener

	// The list of currently opened connections,
	// which will be loop-closed on shutdown
	OpenConnections []net.Conn

	// Port
	Port int

	// Flags for the status of TempoDB
	IsStarting     bool
	IsShuttingDown bool
	IsRunning      bool

	// The configuration
	Cfgs map[string]any

	// Logging
	Logger *slog.Logger

	// This manages telemetry
	Telemetry telemetry.Telemetry

	// Manages write-ahead logging
	Wal wal.Wal
}

func (t Tempo) Start() (err error) {
	t.Net, err = net.Listen("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		fmt.Println("Error starting TempoDB server:", err)
		return fmt.Errorf("failed to start TempoDB server: %w", err)
	}

	fmt.Printf("TempoDB server listening on :%d\n", t.Port)
	ctx := context.Background()
	for {
		conn, _ := t.Net.Accept()
		go handleConn(ctx, conn)

		if t.IsShuttingDown {
			conn.Close()
			break
		}
	}

	return err
}
