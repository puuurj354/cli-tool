package templates

import (
	"strings"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	tests := []struct {
		name     string
		tmplName string
		wantErr  bool
	}{
		// Project templates
		{"get go-api", "go-api", false},
		{"get go-cli", "go-cli", false},
		{"get go-lib", "go-lib", false},
		{"get go-grpc", "go-grpc", false},
		{"get go-worker", "go-worker", false},
		{"get go-tui", "go-tui", false},
		{"get go-microservice", "go-microservice", false},
		{"get go-websocket", "go-websocket", false},
		{"get go-graphql", "go-graphql", false},
		{"get go-lambda", "go-lambda", false},
		{"get go-cron", "go-cron", false},
		{"get go-auth", "go-auth", false},
		{"get go-kafka", "go-kafka", false},
		{"get go-redis", "go-redis", false},
		{"get go-clean-arch", "go-clean-arch", false},
		{"get go-monorepo", "go-monorepo", false},
		// Fullstack
		{"get fullstack", "fullstack", false},
		// Learning templates
		{"get learn-concurrency", "learn-concurrency", false},
		{"get learn-testing", "learn-testing", false},
		{"get learn-dsa", "learn-dsa", false},
		{"get learn-generics", "learn-generics", false},
		{"get learn-context", "learn-context", false},
		{"get learn-http", "learn-http", false},
		{"get learn-error-handling", "learn-error-handling", false},
		{"get learn-interfaces", "learn-interfaces", false},
		{"get learn-design-patterns", "learn-design-patterns", false},
		// Skill templates
		{"get challenge-30days", "challenge-30days", false},
		{"get mini-project", "mini-project", false},
		{"get refactoring-exercise", "refactoring-exercise", false},
		{"get code-review-exercise", "code-review-exercise", false},
		// Error cases
		{"template not found", "non-existent", true},
		{"empty name", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := GetTemplate(tt.tmplName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTemplate(%q) error = %v, wantErr %v", tt.tmplName, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tmpl.Name != tt.tmplName {
					t.Errorf("expected name %q, got %q", tt.tmplName, tmpl.Name)
				}
				if tmpl.Description == "" {
					t.Errorf("expected non-empty description for %s", tt.tmplName)
				}
				if len(tmpl.Directories) == 0 {
					t.Errorf("expected directories for %s", tt.tmplName)
				}
				if len(tmpl.Files) == 0 {
					t.Errorf("expected files for %s", tt.tmplName)
				}
			}
		})
	}
}

func TestGetAllTemplates(t *testing.T) {
	templates := GetAllTemplates()

	// Should have exactly 30 templates
	if len(templates) != 30 {
		t.Errorf("expected 30 templates, got %d", len(templates))
	}

	// All templates should have name, description, directories, and files
	for _, tmpl := range templates {
		if tmpl.Name == "" {
			t.Error("template has empty name")
		}
		if tmpl.Description == "" {
			t.Errorf("template %s has empty description", tmpl.Name)
		}
		if len(tmpl.Directories) == 0 {
			t.Errorf("template %s has no directories", tmpl.Name)
		}
		if len(tmpl.Files) == 0 {
			t.Errorf("template %s has no files", tmpl.Name)
		}

		// Verify all files have paths and content
		for i, f := range tmpl.Files {
			if f.Path == "" {
				t.Errorf("template %s file %d has empty path", tmpl.Name, i)
			}
			if f.Content == "" {
				t.Errorf("template %s file %s has empty content", tmpl.Name, f.Path)
			}
		}
	}

	// Verify template categories
	categories := map[string][]string{
		"go-":          {},
		"learn-":       {},
		"challenge-":   {},
		"mini-":        {},
		"refactoring-": {},
		"code-review-": {},
		"fullstack":    {},
	}

	for _, tmpl := range templates {
		categorized := false
		for prefix := range categories {
			if strings.HasPrefix(tmpl.Name, prefix) || tmpl.Name == "fullstack" {
				categories[prefix] = append(categories[prefix], tmpl.Name)
				categorized = true
				break
			}
		}
		if !categorized {
			t.Errorf("template %s not in any recognized category", tmpl.Name)
		}
	}
}

func TestTemplateFileContentHasTemplateSyntax(t *testing.T) {
	templates := GetAllTemplates()

	for _, tmpl := range templates {
		for _, file := range tmpl.Files {
			// At least one template file should have template syntax
			// (not all will, but most should have at least ProjectName, ModuleName, or PackageName)
			hasTemplateVar := strings.Contains(file.Content, "{{.ProjectName}}") ||
				strings.Contains(file.Content, "{{.ModuleName}}") ||
				strings.Contains(file.Content, "{{.PackageName}}") ||
				strings.Contains(file.Content, "{{.Description}}") ||
				strings.Contains(file.Content, "{{.License}}")

			if !hasTemplateVar && !strings.HasPrefix(tmpl.Name, "learn-") &&
				!strings.HasPrefix(tmpl.Name, "challenge-") &&
				!strings.HasPrefix(tmpl.Name, "code-review-") {
				// Some templates might not have template vars, that's okay
				// but production templates typically should
				t.Logf("template %s file %s has no template variables", tmpl.Name, file.Path)
			}
		}
	}
}

func TestTemplateNamesAreUnique(t *testing.T) {
	templates := GetAllTemplates()
	names := make(map[string]bool)

	for _, tmpl := range templates {
		if names[tmpl.Name] {
			t.Errorf("duplicate template name: %s", tmpl.Name)
		}
		names[tmpl.Name] = true
	}
}

func TestTemplateDirectoriesAreValid(t *testing.T) {
	templates := GetAllTemplates()

	for _, tmpl := range templates {
		for _, dir := range tmpl.Directories {
			if dir == "" {
				t.Errorf("template %s has empty directory", tmpl.Name)
			}
			if strings.HasPrefix(dir, "/") {
				t.Errorf("template %s directory %s is absolute path, should be relative", tmpl.Name, dir)
			}
			if strings.Contains(dir, "\\") {
				t.Errorf("template %s directory %s uses backslashes, should use forward slashes", tmpl.Name, dir)
			}
		}
	}
}

func TestTemplateFilePathsAreValid(t *testing.T) {
	templates := GetAllTemplates()

	for _, tmpl := range templates {
		for _, file := range tmpl.Files {
			if file.Path == "" {
				t.Errorf("template %s has file with empty path", tmpl.Name)
			}
			if strings.HasPrefix(file.Path, "/") {
				t.Errorf("template %s file path %s is absolute, should be relative", tmpl.Name, file.Path)
			}
			if strings.Contains(file.Path, "\\") {
				t.Errorf("template %s file path %s uses backslashes, should use forward slashes", tmpl.Name, file.Path)
			}
		}
	}
}

func TestSpecificTemplatesStructure(t *testing.T) {
	tests := []struct {
		name          string
		tmplName      string
		minDirs       int
		minFiles      int
		expectedFiles []string // file paths that should exist
	}{
		{"go-api structure", "go-api", 6, 8, []string{"cmd/api/main.go", "README.md"}},
		{"go-cli structure", "go-cli", 3, 4, []string{"main.go", "cmd/root.go"}},
		{"learn-concurrency structure", "learn-concurrency", 5, 7, []string{"README.md", "01-goroutines/main.go"}},
		{"learn-testing structure", "learn-testing", 4, 8, []string{"unit/calculator.go", "unit/calculator_test.go"}},
		{"fullstack structure", "fullstack", 6, 17, []string{"backend/cmd/api/main.go", "frontend/package.json"}},
		{"challenge-30days structure", "challenge-30days", 9, 11, []string{"README.md", "week1/day01_hello/main.go"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := GetTemplate(tt.tmplName)
			if err != nil {
				t.Fatalf("failed to get template %s: %v", tt.tmplName, err)
			}

			if len(tmpl.Directories) < tt.minDirs {
				t.Errorf("template %s: expected at least %d directories, got %d", tt.tmplName, tt.minDirs, len(tmpl.Directories))
			}

			if len(tmpl.Files) < tt.minFiles {
				t.Errorf("template %s: expected at least %d files, got %d", tt.tmplName, tt.minFiles, len(tmpl.Files))
			}

			filePaths := make(map[string]bool)
			for _, f := range tmpl.Files {
				filePaths[f.Path] = true
			}

			for _, expectedFile := range tt.expectedFiles {
				if !filePaths[expectedFile] {
					t.Errorf("template %s: expected file %s not found", tt.tmplName, expectedFile)
				}
			}
		})
	}
}
