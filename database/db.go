package database

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

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
	return DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
}

// IsPresent executes a function that will return a boolean value indicating wether a key is present or not in the BadgerDB
func IsPresent(key string) bool {
	if err := DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return false
	}

	return true
}

// Get executes a function that will print the value found with key in BadgerDB
func Get(key string) error {
	return DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		fmt.Println(string(val))

		return nil
	})
}

// Delete executes a function that will delete, if present, the value with specified key in BadgerDB
func Delete(key string) error {
	return DB.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		if err != nil {
			return err
		}

		return nil
	})
}

// CloseDb closes BadgerDB
func CloseDb() {
	DB.Close()
}
