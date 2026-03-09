package generator

import (
	"testing"

	"github.com/purnama/scaffold/internal/templates"
)

func TestAllTemplatesParse(t *testing.T) {
	// Dummy data for substitution
	data := TemplateData{
		ProjectName: "testproject",
		PackageName: "testproject",
		ModuleName:  "github.com/test/testproject",
		Description: "Test Project",
		License:     "MIT",
	}

	allTemplates := templates.GetAllTemplates()

	for _, tmpl := range allTemplates {
		t.Run(tmpl.Name, func(t *testing.T) {
			for _, f := range tmpl.Files {
				// Test processing
				_, err := processTemplate(f.Content, data)
				if err != nil {
					t.Errorf("Failed to process file %s in template %s: %v", f.Path, tmpl.Name, err)
				}

				// Also test path processing for completeness
				_ = processPath(f.Path, data)
			}
		})
	}
}
