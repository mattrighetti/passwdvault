package database

import (
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

// DB instance of the BadgerDB
var DB *badger.DB

// DbInit executes a function that initiates and opens BadgerDB
func DbInit() {
	var err error
	DB, err = badger.Open(badger.DefaultOptions("/tmp/badger").WithLogger(nil).WithEncryptionKey([]byte("this-is-a-masterkey-with-more-16")))
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDb closes BadgerDB
func CloseDb() {
	DB.Close()
}
