package main

import (
	"github.com/MattRighetti/passwdvault/cmd"
	db "github.com/MattRighetti/passwdvault/database"
)

func main() {
	db.DbInit()
	defer db.CloseDb()
	cmd.Execute()
}
