package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get [PASSWORD_ID]",
	Short: "Prints the password",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
