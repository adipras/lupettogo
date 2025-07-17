package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your development environment for LupettoGo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Running environment check...")

		check("Go", "go", "version", "go version", validateGoVersion)

		check("Git", "git", "--version", "git version", nil)

		check("PostgreSQL (default)", "psql", "--version", "psql", nil)

		check("MySQL (optional)", "mysql", "--version", "mysql", nil)

		check("Docker (optional)", "docker", "--version", "docker", nil)

		fmt.Println("✅ Templates embedded in binary")
	},
}

func check(name, bin, arg, expect string, validator func(string) bool) {
	cmd := exec.Command(bin, arg)
	out, err := cmd.CombinedOutput()
	outStr := strings.TrimSpace(string(out))

	if err != nil {
		if strings.Contains(name, "optional") {
			fmt.Printf("⚠️  %s not found (optional)\n", name)
		} else {
			fmt.Printf("❌ %s not found - please install %s\n", name, bin)
		}
		return
	}

	if validator != nil && !validator(outStr) {
		fmt.Printf("⚠️  %s found but version may not be compatible: %s\n", name, outStr)
		fmt.Printf("   Required: Go 1.18 or higher\n")
		return
	}

	fmt.Printf("✅ %s OK: %s\n", name, outStr)
}

func validateGoVersion(output string) bool {
	re := regexp.MustCompile(`go(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) < 3 {
		return false
	}
	
	major, err1 := strconv.Atoi(matches[1])
	minor, err2 := strconv.Atoi(matches[2])
	
	if err1 != nil || err2 != nil {
		return false
	}
	
	return major > 1 || (major == 1 && minor >= 18)
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
