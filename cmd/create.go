package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

const (
	flagIdentifier      = "identifier"
	flagIdentifierShort = "i"
	flagPasswd          = "passwd"
	flagPasswdShort     = "p"
)

var (
	identifier string
	passwd     string

	createCmd = &cobra.Command{
		Use:   "create [PASSWORD_ID] [PASSWORD]",
		Short: "Creates password with identifier",
		Long:  "examples here...",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return configuration.InitCriticalData()
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Writing (%s, %s) to BadgerDB\n", identifier, passwd)
			reader := bufio.NewReader(os.Stdin)

			if db.IsPresent(identifier) {
				fmt.Print("Value already exist, would you want to overwrite it? [y/n]: ")
			} else {
				fmt.Print("Confirm? [y/n]: ")
			}

			text, _ := reader.ReadString('\n')
			text = strings.ReplaceAll(text, "\n", "")

			if text == "y" {
				db.Write(identifier, passwd)
				fmt.Println("Password successfully saved.")
			} else {
				fmt.Println("Aborting")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&identifier, flagIdentifier, flagIdentifierShort, "", "Password identifier")
	createCmd.Flags().StringVarP(&passwd, flagPasswd, flagPasswdShort, "", "Password")
	createCmd.MarkFlagRequired(flagIdentifier)
	createCmd.MarkFlagRequired(flagPasswd)
}
