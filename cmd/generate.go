package cmd

import "github.com/spf13/cobra"

var generateCmd = &cobra.Command{
	Use:   "generate [PASSWORD_ID]",
	Short: "Generates password",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
