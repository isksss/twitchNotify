package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DiscordBotToken string `toml:"discord_bot_token"`
}

// Load reads config from $XDG_CONFIG_HOME/twitchNotify/config.toml
func Load() (*Config, error) {
	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfig == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		xdgConfig = filepath.Join(homeDir, ".config")
	}
	configPath := filepath.Join(xdgConfig, "twitchNotify", "config.toml")

	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, errors.New("failed to decode config file: " + err.Error())
	}
	if cfg.DiscordBotToken == "" {
		return nil, errors.New("discord_bot_token is missing in config file")
	}
	return &cfg, nil
}
