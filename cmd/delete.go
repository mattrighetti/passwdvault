package cmd

import (
	"fmt"
	"os"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [PASSWORD_ID]",
	Short: "Deletes password",
	Long:  "Delete will search for the password of the passed pasword identifier and, if present, it will delete it",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.InitCriticalData()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

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
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.CloseDb()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
