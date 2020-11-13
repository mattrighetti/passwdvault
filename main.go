package main

import (
	"github.com/MattRighetti/passwdvault/cmd"
	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
)

func main() {
	err := configuration.CheckInitFile()
	if err != nil {
		panic("could not find or create configuration file")
	}

	db.DbInit()
	defer db.CloseDb()

	cmd.Execute()
}
