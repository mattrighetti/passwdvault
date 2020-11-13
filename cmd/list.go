package cmd

import (
	"fmt"

	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all password identifiers available",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := db.GetAllKeys()
		if err != nil {
			fmt.Println("Could not print all keys")
		}

		for _, key := range keys {
			fmt.Println(string(key))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
