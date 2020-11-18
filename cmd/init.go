package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/MattRighetti/passwdvault/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	stdinReader = bufio.NewReader(os.Stdin)

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes configuration files and database",
		PreRun: func(cmd *cobra.Command, args []string) {
			exists := utils.FileExists(configuration.ConfigFilePath)
			if exists {
				overwrite, err := utils.ReadBool("A configuration file already exist, would you like to overwrite it?")

				if err != nil {
					log.Fatal(err)
				}

				if overwrite {
					os.Exit(0)
				}
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var config configuration.Configuration
			var masterkey []byte = nil

			unamePattern := `^[a-zA-Z0-9]+(?:[-_][a-zA-Z0-9]+)*$`
			emailPattern := "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

			utils.ReadValidatedInputString("Your name: ", &config.User.Name, unamePattern, 8, 20)
			utils.ReadValidatedInputString("Your email: ", &config.User.Email, emailPattern, 3, 254)

			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			config.Database.Path = dir
			config.Database.Name = ".passwddatabase"

			confirm, err := utils.ReadBool("Would you like to encrypt the database?")
			if err != nil {
				log.Fatal(err)
			}

			if confirm {
				config.Database.Encrypted = true

				for {
					masterkey, _ = utils.ReadInputStringHideInput("MasterKey (must be either 16 or 24 or 64 chars): ")
					if len(masterkey) == 16 || len(masterkey) == 24 || len(masterkey) == 64 {
						break
					}
				}
				config.Database.MasterKey.Length = int8(len(masterkey))
				viper.Set("masterkey", string(masterkey))

				confirm, err := utils.ReadBool("Would you like to store the masterkey somewhere in your system in order to avoid writing everytime?")

				if err != nil {
					log.Fatal(err)
				}

				if confirm {
					fmt.Print("Insert path to file: ")
					fmt.Scanf("%s", &config.Database.MasterKey.FromFilePath)
				}

			} else {
				config.Database.Encrypted = false
			}

			dbPath := path.Join(config.Database.Path, config.Database.Name)
			configuration.CreateDb(dbPath, masterkey)
			configuration.CreateConfigurationFile(&config.User, &config.Database)
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}
