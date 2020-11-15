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
			var masterkey []byte = nil

			utils.ReadInputString("Your name: ", &config.User.Name)
			utils.ReadInputString("Your email: ", &config.User.Email)

			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("saving folder in %s\n", dir)

			config.Database.Path = dir
			config.Database.Name = ".passwddatabase"

			confirm := utils.ReadBool("Would you like to encrypt the database?")
			if confirm {
				config.Database.Encrypted = true

				for {
					masterkey, _ = utils.ReadInputStringHideInput("MasterKey (must be either 16 or 24 or 64 chars): ")
					log.Printf("Read %d bytes\n", len(masterkey))
					if len(masterkey) == 16 || len(masterkey) == 24 || len(masterkey) == 64 {
						break
					}
				}
				config.Database.MasterKey.Length = int8(len(masterkey))
				viper.Set("masterkey", string(masterkey))

				confirm := utils.ReadBool("Would you like to store the masterkey somewhere in your system in order to avoid writing everytime?")
				if confirm {
					fmt.Print("Insert path to file: ")
					fmt.Scanf("%s", &config.Database.MasterKey.FromFilePath)
				}

			} else {
				config.Database.Encrypted = false
			}

			configuration.CreateConfigurationFile(&config.User, &config.Database)
			configuration.DbInit()
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}
