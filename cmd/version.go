package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VERSION = "v1.0"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Password Generator CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Password Generator CLI", VERSION)
	},
}
