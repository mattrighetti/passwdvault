package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "passwdvault [COMMAND]",
		Short: "PasswdVault is a powerfult and secure CLI password manager",
		Long:  "A password manager built with security in mind without sacrificing ease of use.",
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
