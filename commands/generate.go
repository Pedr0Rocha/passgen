package commands

import (
	"fmt"
	"log"

	"github.com/pedr0rocha/password-generator-cli/password"
	"github.com/spf13/cobra"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 64
)

var (
	passwordLength int
	hasSymbols     bool
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&passwordLength, "length", "l", 15, "Password length")
	generateCmd.Flags().BoolVarP(&hasSymbols, "symbols", "s", true, "Enable symbols to the generated password")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates a random password",
	Long:  `Generates a random password mixing symbols, letters and numbers`,
	Args: func(cmd *cobra.Command, args []string) error {
		if passwordLength < MinPasswordLength || passwordLength > MaxPasswordLength {
			return fmt.Errorf("invalid password length, value must be between %d-%d", MinPasswordLength, MaxPasswordLength)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		password, err := password.Generate(passwordLength, hasSymbols)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(passwordLength, "character password generated:")
		fmt.Println(password)
	},
}
