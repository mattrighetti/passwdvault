package database

import (
	"log"
	"path"

	"github.com/MattRighetti/passwdvault/configuration"
	badger "github.com/dgraph-io/badger/v2"
)

// DB instance of the BadgerDB
var DB *badger.DB

// DbInit executes a function that initiates and opens BadgerDB
func DbInit() {
	var err error
	databaseFilePath := path.Join(configuration.DefaultConfig.Database.Path, configuration.DefaultConfig.Database.Name)
	if configuration.DefaultConfig.Database.Encrypted {
		masterKey, err := configuration.ReadMasterKeyFromFile(configuration.DefaultConfig.Database.MasterKey.FromFilePath)
		log.Printf("Read MasterKey: %s\n", masterKey)
		if err != nil {
			log.Fatal(err)
		}

		DB, err = badger.Open(badger.DefaultOptions(databaseFilePath).WithLogger(nil).WithEncryptionKey(masterKey))
	} else {
		DB, err = badger.Open(badger.DefaultOptions(databaseFilePath).WithLogger(nil))
	}

	if err != nil {
		log.Fatal(err)
	}
}

// CloseDb closes BadgerDB
func CloseDb() {
	DB.Close()
}
