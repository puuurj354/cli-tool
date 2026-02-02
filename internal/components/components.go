// =============================================================================
// COMPONENTS PACKAGE
// =============================================================================
// Provides reusable components that can be added to existing projects.
// Each component is a set of files with a specific purpose.

package components

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed embedded/*
var embeddedFS embed.FS

// -----------------------------------------------------------------------------
// Types
// -----------------------------------------------------------------------------

// Component represents a reusable project component.
type Component struct {
	Name        string          // Component identifier (e.g., "dockerfile")
	Description string          // Human-readable description
	Files       []ComponentFile // Files to create
}

// ComponentFile represents a single file within a component.
type ComponentFile struct {
	Path    string // Relative path from project root
	Content string // File content (template)
}

// -----------------------------------------------------------------------------
// Component Registry
// -----------------------------------------------------------------------------

var componentRegistry map[string]Component

func init() {
	initComponentRegistry()
}

// loadEmbedded loads template content from embedded files.
func loadEmbedded(name string) string {
	data, err := embeddedFS.ReadFile("embedded/" + name)
	if err != nil {
		panic(fmt.Sprintf("failed to load embedded component %s: %v", name, err))
	}
	return string(data)
}

// initComponentRegistry initializes all available components.
func initComponentRegistry() {
	componentRegistry = map[string]Component{
		"dockerfile": {
			Name:        "dockerfile",
			Description: "Multi-stage Dockerfile for Go applications",
			Files: []ComponentFile{
				{Path: "Dockerfile", Content: loadEmbedded("add_dockerfile.tmpl")},
			},
		},
		"makefile": {
			Name:        "makefile",
			Description: "Common Makefile targets (build, test, lint, run)",
			Files: []ComponentFile{
				{Path: "Makefile", Content: loadEmbedded("add_makefile.tmpl")},
			},
		},
		"github-actions": {
			Name:        "github-actions",
			Description: "GitHub Actions CI workflow (test, lint, build)",
			Files: []ComponentFile{
				{Path: ".github/workflows/ci.yml", Content: loadEmbedded("add_github_actions.tmpl")},
			},
		},
		"middleware": {
			Name:        "middleware",
			Description: "HTTP middleware collection (logging, CORS, recovery, rate-limit)",
			Files: []ComponentFile{
				{Path: "internal/middleware/logging.go", Content: loadEmbedded("add_middleware_logging.tmpl")},
				{Path: "internal/middleware/cors.go", Content: loadEmbedded("add_middleware_cors.tmpl")},
				{Path: "internal/middleware/recovery.go", Content: loadEmbedded("add_middleware_recovery.tmpl")},
				{Path: "internal/middleware/ratelimit.go", Content: loadEmbedded("add_middleware_ratelimit.tmpl")},
			},
		},
		"config": {
			Name:        "config",
			Description: "Environment-based configuration with defaults",
			Files: []ComponentFile{
				{Path: "internal/config/config.go", Content: loadEmbedded("add_config.tmpl")},
			},
		},
		"docker-compose": {
			Name:        "docker-compose",
			Description: "Docker Compose with PostgreSQL, Redis, and app service",
			Files: []ComponentFile{
				{Path: "docker-compose.yml", Content: loadEmbedded("add_docker_compose.tmpl")},
			},
		},
		"gitignore": {
			Name:        "gitignore",
			Description: "Go-specific .gitignore file",
			Files: []ComponentFile{
				{Path: ".gitignore", Content: loadEmbedded("add_gitignore.tmpl")},
			},
		},
		"readme": {
			Name:        "readme",
			Description: "Project README template with badges and sections",
			Files: []ComponentFile{
				{Path: "README.md", Content: loadEmbedded("add_readme.tmpl")},
			},
		},
	}
}

// -----------------------------------------------------------------------------
// Public API
// -----------------------------------------------------------------------------

// GetComponent retrieves a component by name.
func GetComponent(name string) (Component, bool) {
	comp, found := componentRegistry[name]
	return comp, found
}

// GetAllComponents returns all registered components.
func GetAllComponents() []Component {
	components := make([]Component, 0, len(componentRegistry))
	for _, c := range componentRegistry {
		components = append(components, c)
	}
	return components
}

// AddComponent adds a component's files to the specified directory.
// If force is false, it will not overwrite existing files.
func AddComponent(targetDir, componentName string, force bool) error {
	comp, found := GetComponent(componentName)
	if !found {
		return fmt.Errorf("unknown component: %s", componentName)
	}

	for _, file := range comp.Files {
		targetPath := filepath.Join(targetDir, file.Path)

		// Check if file exists
		if _, err := os.Stat(targetPath); err == nil && !force {
			return fmt.Errorf("file already exists: %s (use --force to overwrite)", file.Path)
		}

		// Create parent directories
		dir := filepath.Dir(targetPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Write file
		if err := os.WriteFile(targetPath, []byte(file.Content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", file.Path, err)
		}
	}

	return nil
}
