package cmd

import (
	"fmt"

	generator "github.com/MattRighetti/passwdgen"
	"github.com/spf13/cobra"
)

var (
	length int16

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates password",
		Long:  "examples here...",
		Run: func(cmd *cobra.Command, args []string) {
			pass, err := generator.Generate(uint8(length))
			if err != nil {
				fmt.Println("Length not supported.")
			}

			fmt.Println(string(pass))
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().Int16VarP(&length, "length", "l", 8, "Length of the password to generate")
}
