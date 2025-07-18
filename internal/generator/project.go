package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

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
	return processTemplateFS(templatesFS, "templates", dest, data)
}

func processTemplateFS(fsys embed.FS, src, dest string, data ProjectData) error {
	entries, err := fsys.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			// Skip modules directory as it's for module generation
			if entry.Name() == "modules" {
				continue
			}

			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}

			if err := processTemplateFS(fsys, srcPath, destPath, data); err != nil {
				return err
			}
			continue
		}

		// Skip files based on configuration
		if shouldSkipFile(entry.Name(), data) {
			continue
		}

		// Read embedded template file
		content, err := fsys.ReadFile(srcPath)
		if err != nil {
			return err
		}

		// Process template
		tmpl, err := template.New("file").Parse(string(content))
		if err != nil {
			return err
		}

		// Create target file
		file, err := os.Create(destPath)
		if err != nil {
			return err
		}

		// Execute template
		err = tmpl.Execute(file, data)
		file.Close()
		if err != nil {
			return err
		}
	}

	return nil
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
