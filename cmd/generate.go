package cmd

import (
	"fmt"
	"log"

	"github.com/pedr0rocha/password-generator-cli/password"
	"github.com/spf13/cobra"
)

/*
	ADD SUPPORT TO
	- uppercase
	- lowercase
	- symbols
	- digits
	- password length
*/

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a random password.",
	Long:  `Generates a random password mixing symbols, letter and numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, err := password.Generate()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(password)
	},
}
