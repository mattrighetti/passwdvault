package cmd

import "github.com/spf13/cobra"

var searchCmd = &cobra.Command{
	Use:   "search [NAME]",
	Short: "Returns available password with pattern [NAME]",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
