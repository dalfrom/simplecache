package tcp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/dalfrom/simplecache/pkg/cache"
	"github.com/dalfrom/simplecache/pkg/config"
	"github.com/dalfrom/simplecache/pkg/wal"
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
	Config config.Config

	// The write-ahead log
	WAL *wal.Wal

	// Logging
	Logger *slog.Logger
}

func (t ServerCache) Start(configPath string) (err error) {
	t.Net, err = net.Listen("tcp", fmt.Sprintf(":%d", t.Port))
	if err != nil {
		t.Logger.Error("Error starting simplecache server:", slog.Any("err", err))
		return fmt.Errorf("failed to start simplecache server: %w", err)
	}

	t.Config, err = config.ExtractConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to extract config: %w. Are you sure the config is correctly set?", err)
	}

	// Create the cache
	sc := cache.CreateCache()

	// Instead of "just creating the WAL", we also check if there is a previous WAL
	// if there is, we have crashed from the previous process, so we recover the whole WAL
	// Then, we replay the WAL to the cache via a different system
	wal, repopulate := wal.RestoreOrCreateAnew(t.Config.Wal.Dir)
	if repopulate {
		// If we are repopulating, we need to replay the WAL
		// This is a placeholder for the actual replay logic
		lines, err := wal.ReplayWal()
		if err != nil {
			return fmt.Errorf("failed to replay WAL: %w", err)
		}

		// Replay the WAL to the cache
		for _, line := range lines {
			handleOperation(nil, line, sc) // nil connection since we are not responding to anyone
		}
	}

	// Lastly the new WAL is set
	// the old one isn't cleared since, if we crash again, we recover again
	// It'll be cleared via the parameters set on its configuration
	t.WAL = wal

	// Run the goroutine for the WAL clearing every second
	go t.WAL.FlushOldEntries(t.Config.Wal.MaxSize, t.Config.Wal.MaxTime)

	// Handle incoming connections
	fmt.Printf("simplecache server listening on :%d\n", t.Port)
	for {
		conn, _ := t.Net.Accept()
		go handleConn(conn, &t, sc)
	}
}

func (t ServerCache) Stop(force bool) {
	// Close the WAL checker
	t.WAL.StopFlush <- true
	t.WAL.Ticker.Stop()

	if !force {
		for _, conn := range t.OpenConnections {
			conn.Close()
		}
	}

	t.Net.Close()
}
