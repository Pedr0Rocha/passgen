package cmd

import (
	"fmt"
	"log"

	"github.com/pedr0rocha/password-generator-cli/password"
	"github.com/spf13/cobra"
)

var PasswordLength int
var HasSymbols bool

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&PasswordLength, "length", "l", 15, "Password length")
	generateCmd.Flags().BoolVarP(&HasSymbols, "symbols", "s", true, "Enable symbols to the generated password")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a random password.",
	Long:  `Generates a random password mixing symbols, letter and numbers.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if PasswordLength > 64 {
			return fmt.Errorf("invalid password length (max. 64)")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		password, err := password.Generate(PasswordLength, HasSymbols)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(password)
	},
}
