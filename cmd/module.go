package cmd

import (
	"github.com/adipras/lupettogo/internal/generator"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "generate module [name]",
	Short: "Generate a new module (handler, service, model, repo)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.GenerateModule(args[0])
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
}
