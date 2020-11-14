package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/spf13/cobra"
)

var (
	stdinReader = bufio.NewReader(os.Stdin)

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes configuration files and database",
		PreRun: func(cmd *cobra.Command, args []string) {
			err := configuration.CheckInitFile()
			if err == nil {
				var res string
				fmt.Print("A configuration file already exist, would you like to overwrite it? [y/n]: ")
				fmt.Scanf("%s", &res)
				if res != "y" {
					os.Exit(0)
				}
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var masterkey []byte
			var encryptString string
			var user configuration.UserConfiguration
			var databaseConf configuration.DatabaseConfiguration

			readFromStdin(&user.Name, "What's your name? ")
			readFromStdin(&user.Email, "What's your email address? ")
			readFromStdin(&databaseConf.Name, "Insert name of the database folder: ")
			readFromStdin(&databaseConf.Path, "Insert path where you would like to save the database folder: ")
			readFromStdin(&encryptString, "Would you like to encrypt the database? [y/n]: ")

			if encryptString != "y" {
				databaseConf.Encrypted = false
			} else {
				databaseConf.Encrypted = true
			}

			if databaseConf.Encrypted {
				readFromStdin(&databaseConf.MasterKey.FromFilePath, "Insert master key file path (leave empty if you don't want to store it locally): ")
				for {
					masterkey, err = readMasterKeyWithDoubleCheck("Insert master key with which you would like to encrypt the database")
					if err == nil {
						break
					}
					fmt.Println(err)
				}
				// databaseConf.MasterKey.Length = len(masterkey)
				fmt.Print(masterkey)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func readFromStdin(val *string, diplayText string) error {
	fmt.Print(diplayText)
	text, err := stdinReader.ReadString('\n')
	if err != nil {
		return err
	}
	text = strings.ReplaceAll(text, "\n", "")

	val = &text
	return nil
}

func readMasterKeyWithDoubleCheck(displayText string) ([]byte, error) {
	fmt.Println(displayText)
	var mk []byte
	nBytes, err := stdinReader.Read(mk)
	if err != nil {
		return nil, err
	}

	fmt.Println("Type again to confirm:")
	var mkCheck []byte
	nBytesCheck, err := stdinReader.Read(mkCheck)
	if err != nil {
		return nil, err
	}

	if nBytes != 8 && nBytes != 16 && nBytes != 32 && nBytes != 64 {
		return nil, fmt.Errorf("masterKey length is incorrect")
	}

	if nBytes != nBytesCheck {
		return nil, fmt.Errorf("masterKey check is not correct")
	}

	for i := range mk {
		if mk[i] != mkCheck[i] {
			return nil, fmt.Errorf("masterKey check is not correct")
		}
	}

	return mk, nil
}
