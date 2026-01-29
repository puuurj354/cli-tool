package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds user configuration
type Config struct {
	Author         string `json:"author"`
	DefaultLicense string `json:"default_license"`
	ModulePrefix   string `json:"module_prefix"`
	AutoGit        bool   `json:"auto_git"`
	AutoInstall    bool   `json:"auto_install"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Author:         "",
		DefaultLicense: "MIT",
		ModulePrefix:   "github.com/user",
		AutoGit:        true,
		AutoInstall:    false,
	}
}

// Load loads configuration from ~/.scaffold/config.json
func Load() *Config {
	cfg := DefaultConfig()

	configPath := getConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg
	}

	json.Unmarshal(data, cfg)
	return cfg
}

// Save saves configuration to ~/.scaffold/config.json
func Save(cfg *Config) error {
	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(getConfigPath(), data, 0644)
}

// GetCustomTemplatesDir returns the custom templates directory path
func GetCustomTemplatesDir() string {
	return filepath.Join(getConfigDir(), "templates")
}

func getConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".scaffold")
}

func getConfigPath() string {
	return filepath.Join(getConfigDir(), "config.json")
}