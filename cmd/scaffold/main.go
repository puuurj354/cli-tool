package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/purnama/scaffold/internal/components"
	"github.com/purnama/scaffold/internal/config"
	"github.com/purnama/scaffold/internal/generator"
	"github.com/purnama/scaffold/internal/templates"
	"github.com/purnama/scaffold/internal/tui"
	"github.com/spf13/cobra"
)

var version = "0.2.0"

// Flags
var (
	dryRun bool
	force  bool
)

// Styles
var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7C3AED"))
	categoryStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#10B981"))
	dimStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "scaffold",
		Short: "Generate project folder structures automatically",
		Long: `Scaffold - A CLI tool to generate project folder structures automatically.

Create new projects quickly with predefined templates for Go APIs, CLIs,
libraries, gRPC services, fullstack apps, learning projects, and skill exercises.

Commands:
  init [template]    Initialize a new project (interactive or with template)
  list               List all available templates (categorized)
  info <template>    Show template details before creating
  config             Show current configuration

Examples:
  scaffold init                        # Interactive mode with prompts
  scaffold init go-api                 # Quick start with template
  scaffold init fullstack --dry-run    # Preview without creating
  scaffold list                        # See all 30 templates
  scaffold info learn-dsa              # Check what files will be created

Templates (30):
  Project:    go-api, go-cli, go-lib, go-grpc, go-worker, go-tui,
              go-microservice, go-websocket, go-graphql, go-lambda, go-cron,
              go-auth, go-kafka, go-redis, go-clean-arch, go-monorepo
  Fullstack:  fullstack (Go + React/Vite/Bun/Tailwind)
  Learning:   learn-concurrency, learn-testing, learn-dsa, learn-generics,
              learn-context, learn-http, learn-error-handling,
              learn-interfaces, learn-design-patterns
  Skill:      challenge-30days, mini-project, refactoring-exercise, code-review-exercise

Config: ~/.scaffold/config.json`,
		Version: version,
	}

	initCmd := &cobra.Command{
		Use:   "init [template]",
		Short: "Initialize a new project",
		Long: `Initialize a new project in the current directory.

If no template is specified, interactive mode will guide you through:
  - Project name
  - Template selection  
  - Dockerfile inclusion
  - License selection (MIT, Apache 2.0, GPL 3.0)
  - Git initialization

Flags:
  --dry-run    Preview what files will be created without creating them
  --force      Overwrite existing directory if it already exists

Available Templates:
  go-api                 REST API with clean architecture
  go-cli                 CLI app with Cobra
  go-lib                 Go library/package
  go-grpc                gRPC service with proto file
  go-worker              Background worker with job queue
  go-tui                 Terminal UI with Bubbletea
  fullstack              Go backend + React/Vite/Bun/Tailwind frontend
  learn-concurrency      Practice Go concurrency patterns
  learn-testing          Practice Go testing techniques
  learn-dsa              Practice Data Structures & Algorithms
  challenge-30days       30-day Go coding challenge
  mini-project           Mini projects to build (todo-cli, url-shortener)
  refactoring-exercise   Practice refactoring bad code
  code-review-exercise   Find bugs in code (code review practice)

Examples:
  scaffold init                        # Interactive mode
  scaffold init go-api                 # Create Go API project
  scaffold init fullstack              # Create fullstack project (auto-installs deps)
  scaffold init learn-dsa              # Practice DSA with tests
  scaffold init go-api --dry-run       # Preview only
  scaffold init go-api --force         # Overwrite if exists`,
		Args: cobra.MaximumNArgs(1),
		RunE: runInit,
	}
	initCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview files without creating them")
	initCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing files")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List available templates",
		RunE:  runList,
	}

	infoCmd := &cobra.Command{
		Use:   "info <template>",
		Short: "Show template details",
		Long:  `Show detailed information about a template including files and directories that will be created.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runInfo,
	}

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Show or edit configuration",
		RunE:  runConfig,
	}

	addCmd := &cobra.Command{
		Use:   "add <component>",
		Short: "Add component to existing project",
		Long: `Add reusable components to an existing project.

Available Components:
  dockerfile       Multi-stage Dockerfile for Go applications
  makefile         Common Makefile targets (build, test, lint, run)
  github-actions   GitHub Actions CI workflow (test, lint, build)
  middleware       HTTP middleware collection (logging, CORS, recovery, rate-limit)
  config           Environment-based configuration with defaults
  docker-compose   Docker Compose with PostgreSQL, Redis, and app
  gitignore        Go-specific .gitignore file
  readme           Project README template with badges and sections

Examples:
  scaffold add dockerfile       # Add multi-stage Dockerfile
  scaffold add middleware       # Add HTTP middleware collection
  scaffold add github-actions   # Add CI workflow
  scaffold add dockerfile --force  # Overwrite existing files`,
		Args: cobra.ExactArgs(1),
		RunE: runAdd,
	}
	addCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing files")

	rootCmd.AddCommand(initCmd, listCmd, infoCmd, configCmd, addCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runInit(cmd *cobra.Command, args []string) error {
	var cfg tui.ProjectConfig

	if len(args) == 1 {
		templateName := args[0]
		tmpl, err := templates.GetTemplate(templateName)
		if err != nil {
			return fmt.Errorf("template '%s' not found. Use 'scaffold list' to see available templates", templateName)
		}

		// Prompt for project name
		var projectName string
		fmt.Print("Project name: ")
		fmt.Scanln(&projectName)
		if projectName == "" {
			projectName = "my-project"
		}

		cfg = tui.ProjectConfig{
			ProjectName:   projectName,
			TemplateName:  tmpl.Name,
			License:       "MIT",
			IncludeDocker: false,
			InitGit:       true,
		}
	} else {
		var err error
		cfg, err = tui.Run()
		if err != nil {
			return err
		}
	}

	// Set generator options
	opts := generator.Options{
		DryRun: dryRun,
		Force:  force,
	}

	return generator.GenerateWithOptions(cfg, opts)
}

func runList(cmd *cobra.Command, args []string) error {
	tmpls := templates.GetAllTemplates()

	// Group templates by category
	categories := map[string][]templates.Template{
		"Project":   {},
		"Learning":  {},
		"Fullstack": {},
		"Skill":     {},
	}

	for _, t := range tmpls {
		switch {
		case strings.HasPrefix(t.Name, "learn-"):
			categories["Learning"] = append(categories["Learning"], t)
		case t.Name == "fullstack":
			categories["Fullstack"] = append(categories["Fullstack"], t)
		case strings.HasPrefix(t.Name, "challenge-") ||
			strings.HasPrefix(t.Name, "mini-") ||
			strings.HasPrefix(t.Name, "refactoring-") ||
			strings.HasPrefix(t.Name, "code-review-"):
			categories["Skill"] = append(categories["Skill"], t)
		default:
			categories["Project"] = append(categories["Project"], t)
		}
	}

	fmt.Println(titleStyle.Render("📦 Available Templates"))
	fmt.Println()

	// Print in order
	order := []string{"Project", "Fullstack", "Learning", "Skill"}
	for _, cat := range order {
		list := categories[cat]
		if len(list) == 0 {
			continue
		}

		// Sort alphabetically
		sort.Slice(list, func(i, j int) bool {
			return list[i].Name < list[j].Name
		})

		fmt.Println(categoryStyle.Render(cat + ":"))
		for _, t := range list {
			fmt.Printf("  %-20s %s\n", t.Name, dimStyle.Render(t.Description))
		}
		fmt.Println()
	}

	fmt.Println(dimStyle.Render("Usage: scaffold init <template>"))
	fmt.Println(dimStyle.Render("       scaffold info <template>  # Show details"))
	return nil
}

func runInfo(cmd *cobra.Command, args []string) error {
	templateName := args[0]
	tmpl, err := templates.GetTemplate(templateName)
	if err != nil {
		return fmt.Errorf("template '%s' not found. Use 'scaffold list' to see available templates", templateName)
	}

	fmt.Println(titleStyle.Render("📋 Template: " + tmpl.Name))
	fmt.Println()
	fmt.Println(tmpl.Description)
	fmt.Println()

	// Show directories
	if len(tmpl.Directories) > 0 {
		fmt.Println(categoryStyle.Render("Directories:"))
		for _, dir := range tmpl.Directories {
			fmt.Printf("  📁 %s/\n", dir)
		}
		fmt.Println()
	}

	// Show files
	if len(tmpl.Files) > 0 {
		fmt.Println(categoryStyle.Render("Files:"))
		for _, f := range tmpl.Files {
			fmt.Printf("  📄 %s\n", f.Path)
		}
		fmt.Println()
	}

	// Show usage
	fmt.Println(categoryStyle.Render("Usage:"))
	fmt.Printf("  scaffold init %s\n", tmpl.Name)
	fmt.Printf("  scaffold init %s --dry-run  # Preview first\n", tmpl.Name)

	return nil
}

func runConfig(cmd *cobra.Command, args []string) error {
	cfg := config.Load()

	fmt.Println(titleStyle.Render("⚙️  Configuration"))
	fmt.Println()
	fmt.Printf("  Author:         %s\n", valueOrDefault(cfg.Author, "(not set)"))
	fmt.Printf("  Default License: %s\n", cfg.DefaultLicense)
	fmt.Printf("  Module Prefix:  %s\n", cfg.ModulePrefix)
	fmt.Printf("  Auto Git:       %v\n", cfg.AutoGit)
	fmt.Printf("  Auto Install:   %v\n", cfg.AutoInstall)
	fmt.Println()
	fmt.Println(dimStyle.Render("Config file: ~/.scaffold/config.json"))
	fmt.Println(dimStyle.Render("Custom templates: ~/.scaffold/templates/"))

	return nil
}

func getProjectNameFromCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "my-project"
	}
	return filepath.Base(cwd)
}

func valueOrDefault(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

// runAdd handles the add command for adding components to existing projects
func runAdd(cmd *cobra.Command, args []string) error {
	componentName := args[0]

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Check if component exists
	comp, found := components.GetComponent(componentName)
	if !found {
		fmt.Printf("Error: Unknown component '%s'\n\n", componentName)
		fmt.Println("Available components:")
		for _, c := range components.GetAllComponents() {
			fmt.Printf("  %-15s  %s\n", c.Name, c.Description)
		}
		return fmt.Errorf("component not found")
	}

	// Add the component
	fmt.Printf("Adding component: %s\n", titleStyle.Render(comp.Name))
	fmt.Printf("Description: %s\n\n", comp.Description)

	err = components.AddComponent(cwd, componentName, force)
	if err != nil {
		return err
	}

	// Show created files
	fmt.Println(categoryStyle.Render("Created files:"))
	for _, f := range comp.Files {
		fmt.Printf("  ✓ %s\n", f.Path)
	}
	fmt.Println()
	fmt.Println(dimStyle.Render("Tip: Review TODO comments in generated files for customization"))

	return nil
}
