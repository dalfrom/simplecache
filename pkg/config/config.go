package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type User string

type Wal struct {
	Dir     string
	MaxSize int64
	MaxTime int64
}

type Config struct {
	Port int
	Wal  Wal

	ActiveUsers []User

	Users map[User]struct {
		Username string
		Password string
	}
}

// ExtractConfig reads the configuration file and populates the Config struct
// The configFilePath is the path to the TOML file housing the configuration for the system together with the user's credentials
func ExtractConfig(configFilePath string) (cfg Config, err error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("error decoding config file: %w", err)
	}

	return
}
