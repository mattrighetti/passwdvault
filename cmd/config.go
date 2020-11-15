package cmd

import (
	"log"

	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and show PasswdVault configuration",
	Long:  `examples here...`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.CheckForConfigFileAndParse()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("Adding %s.%s to config file", args[0], args[1])
		return addConfigToConfigFile(args[0], args[1])
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
