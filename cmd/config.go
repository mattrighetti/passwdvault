package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and show PasswdVault configuration",
	Long:  `examples here...`,
	Run: func(cmd *cobra.Command, args []string) {
		// Config stuff here
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
