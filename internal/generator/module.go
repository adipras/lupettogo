package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ModuleData struct {
	ProjectName string
	ModuleName  string
	ModuleTitle string
}

func GenerateModule(moduleName string) error {
	if moduleName == "" {
		return fmt.Errorf("module name cannot be empty")
	}

	// Get current project name from go.mod
	projectName, err := getCurrentProjectName()
	if err != nil {
		return fmt.Errorf("failed to detect project name: %w", err)
	}

	data := ModuleData{
		ProjectName: projectName,
		ModuleName:  strings.ToLower(moduleName),
		ModuleTitle: strings.Title(moduleName),
	}

	if err := generateModuleFiles(data); err != nil {
		return fmt.Errorf("failed to generate module files: %w", err)
	}

	fmt.Printf("✅ Module '%s' created successfully!\n", moduleName)
	fmt.Printf("📝 Don't forget to:\n")
	fmt.Printf("   - Add the new model to database migrations\n")
	fmt.Printf("   - Register the handler in server routes\n")
	fmt.Printf("   - Update services.go and handlers.go\n")
	return nil
}

func generateModuleFiles(data ModuleData) error {
	templateDir := "templates/modules"
	
	// Check if we're in a LupettoGo CLI directory or generated project
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return fmt.Errorf("module templates not found. Make sure you're running this from the LupettoGo CLI directory")
	}

	files := map[string]string{
		"model.go.tmpl":      fmt.Sprintf("internal/models/%s.go", data.ModuleName),
		"repository.go.tmpl": fmt.Sprintf("internal/repositories/%s_repository.go", data.ModuleName),
		"service.go.tmpl":    fmt.Sprintf("internal/services/%s_service.go", data.ModuleName),
		"handler.go.tmpl":    fmt.Sprintf("internal/handlers/%s_handler.go", data.ModuleName),
	}

	for templateFile, outputFile := range files {
		templatePath := filepath.Join(templateDir, templateFile)
		
		// Create output directory if it doesn't exist
		outputDir := filepath.Dir(outputFile)
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return err
		}

		// Read and process template
		content, err := os.ReadFile(templatePath)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", templateFile, err)
		}

		// Replace placeholders
		processed := strings.ReplaceAll(string(content), "__module__", data.ModuleName)
		processed = strings.ReplaceAll(processed, "__Module__", data.ModuleTitle)
		processed = strings.ReplaceAll(processed, "{{.ProjectName}}", data.ProjectName)

		// Write output file
		if err := os.WriteFile(outputFile, []byte(processed), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", outputFile, err)
		}

		fmt.Printf("📄 Created %s\n", outputFile)
	}

	return nil
}

func getCurrentProjectName() (string, error) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	return "", fmt.Errorf("module declaration not found in go.mod")
}
