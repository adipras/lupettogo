package cmd

import (
	"fmt"

	"github.com/adipras/lupettogo/internal/generator"
	"github.com/spf13/cobra"
)

var (
	dbDriver   string
	withAuth   bool
	withDocker bool
	withTests  bool
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Generate a new Golang SaaS starter project",
	Long: `Generate a new production-ready Golang SaaS starter project with clean architecture.

The project includes:
- Clean architecture structure (handlers, services, repositories, models)
- Configuration management with Viper
- Database integration with GORM
- HTTP server with Gin
- Middleware support (CORS, logging, recovery)
- Environment configuration
- Docker support (optional)
- Testing infrastructure (optional)

Examples:
  lupettogo init my-saas-app
  lupettogo init my-api --db postgres --with-auth --with-docker
  lupettogo init simple-api --db mysql --with-tests`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		// Validate project name
		if err := validateProjectName(projectName); err != nil {
			return err
		}

		// Show configuration
		fmt.Printf("🐺 Creating project '%s' with:\n", projectName)
		fmt.Printf("   Database: %s\n", dbDriver)
		fmt.Printf("   Auth: %v\n", withAuth)
		fmt.Printf("   Docker: %v\n", withDocker)
		fmt.Printf("   Tests: %v\n", withTests)
		fmt.Println()

		config := generator.ProjectConfig{
			Name:       projectName,
			DBDriver:   dbDriver,
			WithAuth:   withAuth,
			WithDocker: withDocker,
			WithTests:  withTests,
		}

		return generator.GenerateProjectWithConfig(config)
	},
}

func validateProjectName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("project name cannot be empty")
	}
	if len(name) > 50 {
		return fmt.Errorf("project name too long (max 50 characters)")
	}
	// Add more validation as needed
	return nil
}

func init() {
	initCmd.Flags().StringVar(&dbDriver, "db", "postgres", "Database driver (postgres, mysql)")
	initCmd.Flags().BoolVar(&withAuth, "with-auth", false, "Include authentication scaffolding")
	initCmd.Flags().BoolVar(&withDocker, "with-docker", true, "Include Docker configuration")
	initCmd.Flags().BoolVar(&withTests, "with-tests", true, "Include testing infrastructure")

	rootCmd.AddCommand(initCmd)
}
