package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [PASSWORD_ID] [PASSWORD]",
	Short: "Creates password with identifier",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
