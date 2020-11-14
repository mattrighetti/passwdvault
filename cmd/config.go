package cmd

import (
	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and show PasswdVault configuration",
	Long:  `examples here...`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := configuration.CheckInitFile()
		if err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Config stuff here
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
