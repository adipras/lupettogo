package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of LupettoGo CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("LupettoGo CLI %s 🐺\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
