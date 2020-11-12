package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete [PASSWORD_ID]",
	Short: "Deletes password",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
