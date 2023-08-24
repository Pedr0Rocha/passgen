package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	VERSION string = "v1.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Password Generator CLI",
	Run:   printVersion,
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println("passgen", VERSION)
}
