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
	src := "templates"
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

	if err := processTemplateDir(src, dest, data); err != nil {
		return fmt.Errorf("failed to process templates: %w", err)
	}

	fmt.Printf("✅ Project '%s' created successfully!\n", config.Name)
	fmt.Printf("📁 Run 'cd %s && go mod tidy' to get started\n", config.Name)
	return nil
}

func processTemplateDir(src, dest string, data ProjectData) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip modules directory as it's for module generation
		if strings.Contains(path, "modules") {
			return nil
		}

		relPath, _ := filepath.Rel(src, path)
		
		// Skip files based on configuration
		if shouldSkipFile(relPath, data) {
			return nil
		}

		targetPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// Read template file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Process template
		tmpl, err := template.New("file").Parse(string(content))
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
	})
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
