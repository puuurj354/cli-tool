// =============================================================================
// COMPONENTS TEST - TDD First
// =============================================================================
// Tests for the component registry and add functionality.
// These tests are written BEFORE the implementation (Red-Green-Refactor).

package components

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// -----------------------------------------------------------------------------
// Test: GetComponent retrieves a component by name
// -----------------------------------------------------------------------------
func TestGetComponent(t *testing.T) {
	tests := []struct {
		name      string // Name of the component to retrieve
		wantFound bool   // Whether component should be found
	}{
		{"dockerfile", true},
		{"makefile", true},
		{"github-actions", true},
		{"middleware", true},
		{"config", true},
		{"docker-compose", true},
		{"gitignore", true},
		{"readme", true},
		{"nonexistent", false}, // Should not exist
		{"", false},            // Empty name
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp, found := GetComponent(tt.name)
			if found != tt.wantFound {
				t.Errorf("GetComponent(%q) found = %v, want %v", tt.name, found, tt.wantFound)
			}
			if found && comp.Name != tt.name {
				t.Errorf("GetComponent(%q) returned component with name %q", tt.name, comp.Name)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Test: GetAllComponents returns all registered components
// -----------------------------------------------------------------------------
func TestGetAllComponents(t *testing.T) {
	components := GetAllComponents()

	// We expect at least 8 components
	minExpected := 8
	if len(components) < minExpected {
		t.Errorf("GetAllComponents() returned %d components, want at least %d", len(components), minExpected)
	}

	// Verify required components exist
	required := []string{"dockerfile", "makefile", "github-actions", "middleware", "config", "docker-compose", "gitignore", "readme"}
	for _, name := range required {
		found := false
		for _, c := range components {
			if c.Name == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Required component %q not found in GetAllComponents()", name)
		}
	}
}

// -----------------------------------------------------------------------------
// Test: Each component has valid, non-empty files
// -----------------------------------------------------------------------------
func TestComponentHasValidFiles(t *testing.T) {
	components := GetAllComponents()

	for _, comp := range components {
		t.Run(comp.Name, func(t *testing.T) {
			// Each component must have at least one file
			if len(comp.Files) == 0 {
				t.Errorf("Component %q has no files", comp.Name)
			}

			// Each file must have a path and non-empty content
			for _, f := range comp.Files {
				if f.Path == "" {
					t.Errorf("Component %q has a file with empty path", comp.Name)
				}
				if strings.TrimSpace(f.Content) == "" {
					t.Errorf("Component %q file %q has empty content", comp.Name, f.Path)
				}
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Test: Component has description
// -----------------------------------------------------------------------------
func TestComponentHasDescription(t *testing.T) {
	components := GetAllComponents()

	for _, comp := range components {
		if comp.Description == "" {
			t.Errorf("Component %q has no description", comp.Name)
		}
	}
}

// -----------------------------------------------------------------------------
// Test: AddComponent creates files correctly
// -----------------------------------------------------------------------------
func TestAddComponentCreatesFiles(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Add the dockerfile component
	err = AddComponent(tmpDir, "dockerfile", false)
	if err != nil {
		t.Fatalf("AddComponent() error = %v", err)
	}

	// Check that Dockerfile was created
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		t.Error("Dockerfile was not created")
	}

	// Check content is not empty
	content, _ := os.ReadFile(dockerfilePath)
	if len(content) == 0 {
		t.Error("Dockerfile is empty")
	}
}

// -----------------------------------------------------------------------------
// Test: AddComponent does not overwrite existing files without force
// -----------------------------------------------------------------------------
func TestAddComponentNoOverwrite(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create existing file
	existingContent := "EXISTING CONTENT"
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	os.WriteFile(dockerfilePath, []byte(existingContent), 0644)

	// Try to add dockerfile component without force
	err = AddComponent(tmpDir, "dockerfile", false)
	if err == nil {
		t.Error("AddComponent() should return error when file exists and force=false")
	}

	// Verify original content is preserved
	content, _ := os.ReadFile(dockerfilePath)
	if string(content) != existingContent {
		t.Error("Original file content was overwritten")
	}
}

// -----------------------------------------------------------------------------
// Test: AddComponent with force overwrites existing files
// -----------------------------------------------------------------------------
func TestAddComponentForce(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create existing file
	existingContent := "EXISTING CONTENT"
	dockerfilePath := filepath.Join(tmpDir, "Dockerfile")
	os.WriteFile(dockerfilePath, []byte(existingContent), 0644)

	// Add dockerfile component WITH force
	err = AddComponent(tmpDir, "dockerfile", true)
	if err != nil {
		t.Fatalf("AddComponent() with force error = %v", err)
	}

	// Verify content was overwritten
	content, _ := os.ReadFile(dockerfilePath)
	if string(content) == existingContent {
		t.Error("File was not overwritten with force=true")
	}
}

// -----------------------------------------------------------------------------
// Test: AddComponent returns error for unknown component
// -----------------------------------------------------------------------------
func TestAddComponentUnknown(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	err = AddComponent(tmpDir, "unknown-component", false)
	if err == nil {
		t.Error("AddComponent() should return error for unknown component")
	}
}

// -----------------------------------------------------------------------------
// Test: AddComponent creates necessary directories
// -----------------------------------------------------------------------------
func TestAddComponentCreatesDirectories(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Add middleware component (which creates internal/middleware/)
	err = AddComponent(tmpDir, "middleware", false)
	if err != nil {
		t.Fatalf("AddComponent() error = %v", err)
	}

	// Check that internal/middleware directory was created
	middlewareDir := filepath.Join(tmpDir, "internal", "middleware")
	if _, err := os.Stat(middlewareDir); os.IsNotExist(err) {
		t.Error("internal/middleware directory was not created")
	}
}

// -----------------------------------------------------------------------------
// Benchmark: GetComponent performance
// -----------------------------------------------------------------------------
func BenchmarkGetComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetComponent("dockerfile")
	}
}
