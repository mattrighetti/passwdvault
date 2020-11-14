package cmd

import (
	"bufio"
	"fmt"
	"log"
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
			var config configuration.Configuration

			fmt.Print("Your name: ")
			fmt.Scanf("%s", &config.User.Name)

			fmt.Print("Your email: ")
			fmt.Scanf("%s", &config.User.Email)

			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("saving folder in %s\n", dir)

			config.Database.Path = dir
			config.Database.Name = ".passwddatabase"

			var res string
			fmt.Print("Would you like to encrypt the database? [y/n]: ")
			fmt.Scanf("%s", &res)
			if res == "y" {
				config.Database.Encrypted = true

				var masterkey string
				for {
					fmt.Print("MasterKey (must be either 8 or 16 or 32 or 64 chars): ")
					fmt.Scanf("%s", &masterkey)
					log.Printf("Read %d bytes\n", len(masterkey))
					if len(masterkey) == 8 || len(masterkey) == 16 || len(masterkey) == 32 || len(masterkey) == 64 {
						break
					}
				}
				config.Database.MasterKey.Length = int8(len(masterkey))

				fmt.Print("Would you like to store the masterkey somewhere in your system in order to avoid writing everytime? [y/n]: ")
				fmt.Scanf("%s", &res)
				if res == "y" {
					fmt.Print("Insert path to file: ")
					fmt.Scanf("%s", &config.Database.MasterKey.FromFilePath)
				}
				resetString(&res)

			} else {
				config.Database.Encrypted = false
			}
			resetString(&res)

			configuration.CreateConfigurationFile(&config.User, &config.Database)
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

func resetString(s *string) {
	*s = ""
}
