package cmd

import (
	"fmt"
	"os"

	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Sets PasswdVault configuration values",
	Long:  "Set will add the passed configuration values to the .passwdvaultconfig file",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.CheckForConfigFileAndParse()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			os.Exit(0)
		}

		err := addConfigToConfigFile(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func addConfigToConfigFile(key string, value string) error {
	viper.Set(key, value)
	err := configuration.SaveConfigurationFile()
	if err != nil {
		return err
	}

	return nil
}
