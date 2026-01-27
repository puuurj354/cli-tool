package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	if cfg.DefaultLicense != "MIT" {
		t.Errorf("DefaultLicense = %q, want MIT", cfg.DefaultLicense)
	}

	if cfg.ModulePrefix != "github.com/user" {
		t.Errorf("ModulePrefix = %q, want github.com/user", cfg.ModulePrefix)
	}

	if cfg.AutoGit != true {
		t.Errorf("AutoGit = %v, want true", cfg.AutoGit)
	}

	if cfg.AutoInstall != false {
		t.Errorf("AutoInstall = %v, want false", cfg.AutoInstall)
	}
}

func TestLoadDefaultWhenFileNotExists(t *testing.T) {
	// Set HOME to temp directory so config file doesn't exist
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	cfg := Load()

	if cfg == nil {
		t.Fatal("Load returned nil")
	}

	if cfg.DefaultLicense != "MIT" {
		t.Errorf("DefaultLicense = %q, want MIT", cfg.DefaultLicense)
	}
}

func TestSaveAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	// Create config
	cfg := &Config{
		Author:         "John Doe",
		DefaultLicense: "Apache 2.0",
		ModulePrefix:   "github.com/johndoe",
		AutoGit:        true,
		AutoInstall:    true,
	}

	// Save config
	if err := Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify config file was created
	configPath := filepath.Join(tmpDir, ".scaffold", "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("config file not created at %s", configPath)
	}

	// Load config
	loadedCfg := Load()

	if loadedCfg.Author != "John Doe" {
		t.Errorf("Author = %q, want John Doe", loadedCfg.Author)
	}

	if loadedCfg.DefaultLicense != "Apache 2.0" {
		t.Errorf("DefaultLicense = %q, want Apache 2.0", loadedCfg.DefaultLicense)
	}

	if loadedCfg.ModulePrefix != "github.com/johndoe" {
		t.Errorf("ModulePrefix = %q, want github.com/johndoe", loadedCfg.ModulePrefix)
	}

	if loadedCfg.AutoGit != true {
		t.Errorf("AutoGit = %v, want true", loadedCfg.AutoGit)
	}

	if loadedCfg.AutoInstall != true {
		t.Errorf("AutoInstall = %v, want true", loadedCfg.AutoInstall)
	}
}

func TestGetCustomTemplatesDir(t *testing.T) {
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	dir := GetCustomTemplatesDir()

	expectedDir := filepath.Join(tmpDir, ".scaffold", "templates")
	if dir != expectedDir {
		t.Errorf("GetCustomTemplatesDir = %s, want %s", dir, expectedDir)
	}
}

func TestConfigFileStructure(t *testing.T) {
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	cfg := &Config{
		Author:         "Test Author",
		DefaultLicense: "MIT",
		ModulePrefix:   "github.com/test",
		AutoGit:        true,
		AutoInstall:    false,
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify .scaffold directory was created
	scaffoldDir := filepath.Join(tmpDir, ".scaffold")
	if _, err := os.Stat(scaffoldDir); os.IsNotExist(err) {
		t.Errorf(".scaffold directory not created")
	}

	// Verify config.json was created
	configPath := filepath.Join(scaffoldDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("config.json not created")
	}
}

func TestPartialConfigLoad(t *testing.T) {
	tmpDir := t.TempDir()
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	// Save only partial config
	cfg := &Config{
		Author:         "Partial Author",
		DefaultLicense: "GPL 3.0",
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Load config - should use defaults for missing fields
	loadedCfg := Load()

	if loadedCfg.Author != "Partial Author" {
		t.Errorf("Author = %q, want Partial Author", loadedCfg.Author)
	}

	if loadedCfg.DefaultLicense != "GPL 3.0" {
		t.Errorf("DefaultLicense = %q, want GPL 3.0", loadedCfg.DefaultLicense)
	}
}
