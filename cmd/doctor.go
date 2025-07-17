package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your development environment for LupettoGo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Running environment check...")

		check("Go", "go", "version", "go version", func(out string) bool {
			return strings.Contains(out, "go1.18") || strings.Compare(out, "go1.18") > 0
		})

		check("Git", "git", "--version", "git version", nil)

		check("MySQL (optional)", "mysql", "--version", "mysql", nil)

		fmt.Println("✅ Templates embedded in binary")
	},
}

func check(name, bin, arg, expect string, validator func(string) bool) {
	cmd := exec.Command(bin, arg)
	out, err := cmd.CombinedOutput()
	outStr := strings.TrimSpace(string(out))

	if err != nil {
		fmt.Printf("❌ %s not found (%s)\n", name, bin)
		return
	}

	if validator != nil && !validator(outStr) {
		fmt.Printf("⚠️  %s found but version is not compatible: %s\n", name, outStr)
		return
	}

	fmt.Printf("✅ %s OK: %s\n", name, outStr)
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
