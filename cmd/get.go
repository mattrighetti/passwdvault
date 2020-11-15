package cmd

import (
	"fmt"
	"os"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [PASSWORD_ID]",
	Short: "Gets the password of given identifier if present",
	Long:  "Get will search for the password of the passed pasword identifier and, if present, it will print it",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.InitCriticalData()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

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
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.CloseDb()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
