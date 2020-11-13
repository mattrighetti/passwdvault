package cmd

import (
	"fmt"

	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var (
	passwdOnly bool

	getCmd = &cobra.Command{
		Use:   "get [PASSWORD_ID]",
		Short: "Prints the password",
		Long:  "examples here...",
		Run: func(cmd *cobra.Command, args []string) {
			if db.IsPresent(args[0]) {
				val, err := db.Get(args[0])
				if err != nil {
					fmt.Printf("Could not get value with key %s\n", args[0])
				}

				if passwdOnly {
					fmt.Printf("%s\n", string(val))
				} else {
					fmt.Printf("ID: %s -> Password: %s\n", args[0], string(val))
				}
			} else {
				fmt.Printf("No password stored with key %s\n", args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&passwdOnly, "all", "a", true, "Flag to indicated to print passwork only or id and password")
}
