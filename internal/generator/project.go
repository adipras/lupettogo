package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


type ProjectData struct {
	ProjectName string
	DBDriver    string
	WithAuth    bool
	WithDocker  bool
	WithTests   bool
}

type ProjectConfig struct {
	Name       string
	DBDriver   string
	WithAuth   bool
	WithDocker bool
	WithTests  bool
}

func GenerateProject(projectName string) error {
	config := ProjectConfig{
		Name:       projectName,
		DBDriver:   "postgres",
		WithAuth:   false,
		WithDocker: true,
		WithTests:  true,
	}
	return GenerateProjectWithConfig(config)
}

func GenerateProjectWithConfig(config ProjectConfig) error {
	dest := config.Name

	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create project folder: %w", err)
	}

	data := ProjectData{
		ProjectName: config.Name,
		DBDriver:    config.DBDriver,
		WithAuth:    config.WithAuth,
		WithDocker:  config.WithDocker,
		WithTests:   config.WithTests,
	}

	if err := processEmbeddedTemplates(dest, data); err != nil {
		return fmt.Errorf("failed to process templates: %w", err)
	}

	fmt.Printf("✅ Project '%s' created successfully!\n", config.Name)
	fmt.Printf("📁 Run 'cd %s && go mod tidy' to get started\n", config.Name)
	return nil
}

func processEmbeddedTemplates(dest string, data ProjectData) error {
	// Process root level templates
	for filename, content := range templateFiles {
		if shouldSkipFile(filename, data) {
			continue
		}
		if err := processTemplate(dest, filename, content, data); err != nil {
			return err
		}
	}

	// Process internal templates
	for filename, content := range internalTemplates {
		if shouldSkipFile(filename, data) {
			continue
		}
		if err := processTemplate(dest, filename, content, data); err != nil {
			return err
		}
	}

	// Process test templates if enabled
	if data.WithTests {
		for filename, content := range testTemplates {
			if err := processTemplate(dest, filename, content, data); err != nil {
				return err
			}
		}
	}

	return nil
}

func processTemplate(dest, filename, content string, data ProjectData) error {
	targetPath := filepath.Join(dest, filename)
	
	// Create directory if it doesn't exist
	dir := filepath.Dir(targetPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Process template
	tmpl, err := template.New("file").Parse(content)
	if err != nil {
		return err
	}

	// Create target file
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute template
	return tmpl.Execute(file, data)
}

func shouldSkipFile(relPath string, data ProjectData) bool {
	// Skip Docker files if Docker is disabled
	if !data.WithDocker && (strings.Contains(relPath, "Dockerfile") || strings.Contains(relPath, "docker-compose")) {
		return true
	}

	// Skip test files if tests are disabled
	if !data.WithTests && strings.Contains(relPath, "_test.go") {
		return true
	}

	// Skip auth files if auth is disabled (when implemented)
	if !data.WithAuth && strings.Contains(relPath, "auth") {
		return true
	}

	return false
}
