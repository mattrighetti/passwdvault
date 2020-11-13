package main

import (
	"github.com/MattRighetti/passwdvault/cmd"
	"github.com/MattRighetti/passwdvault/configuration"
	db "github.com/MattRighetti/passwdvault/database"
)

func main() {
	configuration.CheckInitFile()
	db.DbInit()
	defer db.CloseDb()
	cmd.Execute()
}
