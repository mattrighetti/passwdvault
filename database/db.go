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
	DB, err = badger.Open(badger.DefaultOptions("/tmp/badger").WithEncryptionKey([]byte("this-is-a-masterkey-with-more-16")))
	if err != nil {
		log.Fatal(err)
	}
}

// Write executes a function that will write (key, value) to BadgerDB
func Write(key string, value string) error {
	txn := DB.NewTransaction(true)
	defer txn.Discard()

	if err := txn.Set([]byte(key), []byte(value)); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// Get executes a function that will print the value found with key in BadgerDB
func Get(key string) ([]byte, error) {
	txn := DB.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	val, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// IsPresent executes a function that will return a boolean value indicating wether a key is present or not in the BadgerDB
func IsPresent(key string) bool {
	txn := DB.NewTransaction(true)
	defer txn.Discard()

	_, err := txn.Get([]byte(key))
	if err != nil {
		return false
	}

	return true
}

// Delete executes a function that will delete, if present, the value with specified key in BadgerDB
func Delete(key string) error {
	txn := DB.NewTransaction(true)
	defer txn.Discard()

	err := txn.Delete([]byte(key))
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

// CloseDb closes BadgerDB
func CloseDb() {
	DB.Close()
}
