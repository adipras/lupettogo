/*
Copyright © 2025 Adi Prasetyo <adipras2310@gmail.com>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Generate a new Golang SaaS starter project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		src := "templates"
		dest := projectName

		// Buat folder tujuan
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create project folder: %w", err)
		}

		// Salin file dari templates ke folder tujuan
		err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relPath, _ := filepath.Rel(src, path)
			targetPath := filepath.Join(dest, relPath)

			if info.IsDir() {
				return os.MkdirAll(targetPath, os.ModePerm)
			}

			return copyFile(path, targetPath)
		})

		if err != nil {
			return fmt.Errorf("failed to copy template: %w", err)
		}

		fmt.Printf("✅ Project '%s' created successfully!\n", projectName)
		return nil
	},
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func init() {
	rootCmd.AddCommand(initCmd)
}
