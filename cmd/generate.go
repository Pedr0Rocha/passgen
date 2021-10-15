package cmd

import (
	"fmt"
	"log"

	"github.com/pedr0rocha/password-generator-cli/password"
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
		/* @TODO get args
		- password length
		- has symbols
		- has uppercase
		*/

		password, err := password.Generate(15, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(password)
	},
}
