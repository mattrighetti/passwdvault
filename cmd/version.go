package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows passwdvault version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("passwdvault version %s\n", "0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
