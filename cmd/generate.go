package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a random password.",
	Long:  `Generates a random password mixing symbols, letter and numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aaa")
		// @TODO: implement password generator
	},
}
