package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "passgen",
		Short: "A simple password generator.",
		Long: `Password Generator is a CLI library made in Go.
This application is a tool to generate random passwords.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
