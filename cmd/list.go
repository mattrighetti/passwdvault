package cmd

import (
	"fmt"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all password identifiers available",
	Long:  "examples here...",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.InitCriticalData()
	},
	Run: func(cmd *cobra.Command, args []string) {
		keys := db.GetAllKeys()

		for _, key := range keys {
			fmt.Println(string(key))
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.CloseDb()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
