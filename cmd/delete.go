package cmd

import (
	"fmt"

	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [PASSWORD_ID]",
	Short: "Deletes password",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		if db.IsPresent(args[0]) {
			err := db.Delete(args[0])
			if err != nil {
				fmt.Printf("Could not delete value with key %s\n", args[0])
			}

			fmt.Println("Successfully deleted password.")
		} else {
			fmt.Printf("No password stored with key %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
