package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:          "search [PATTERN]",
	Short:        "Returns available password identifiers with that match passed pattern",
	Long:         "Search is an utility tool that will output all your stored password identifiers that match the passed pattern",
	SilenceUsage: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return configuration.InitCriticalData()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		keys := db.GetAllKeys()

		filteredKeys, err := filter(keys, func(in []byte) bool { return bytes.Contains(in, []byte(args[0])) })
		if err != nil {
			fmt.Println(err)
		}

		for _, fKey := range filteredKeys {
			fmt.Println(string(fKey))
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		configuration.CloseDb()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func filter(bytes [][]byte, check func(in []byte) bool) (ret [][]byte, err error) {
	for _, val := range bytes {
		if check(val) {
			ret = append(ret, val)
		}
	}

	if len(ret) == 0 {
		err = fmt.Errorf("no match")
	} else {
		err = nil
	}

	return
}
