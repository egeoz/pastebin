package config

import (
	"log/slog"
	"os"
	"pastebin/src/logger"

	"github.com/pelletier/go-toml/v2"
)

const configFile = "config.toml"

type Config struct {
	Port         int
	Host         string
	Database     string
	DatabaseType string
	LogLevel     string
}

func LoadConfig() Config {
	var cfg Config
	filePath := configFile

	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	configFileContent := string(content)

	err = toml.Unmarshal([]byte(configFileContent), &cfg)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	logger.InitLogger(cfg.LogLevel)

	return cfg
}
