package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "passwdvault [COMMAND]",
		Short: "PasswdVault is a powerfult and secure CLI password manager",
		Long: `		
		PASSWDVAULT

A password manager built with security in mind without sacrificing ease of use.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is the root command")
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// init configuration here
}
