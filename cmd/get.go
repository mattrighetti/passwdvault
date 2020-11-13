package cmd

import (
	"fmt"

	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [PASSWORD_ID]",
	Short: "Prints the password",
	Long:  "examples here...",
	Run: func(cmd *cobra.Command, args []string) {
		if db.IsPresent(args[0]) {
			passwd, err := db.Get(args[0])
			if err != nil {
				fmt.Printf("Could not get value with key %s\n", args[0])
			}

			fmt.Println(string(passwd))
		} else {
			fmt.Printf("No password stored with key %s\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
