package database

import (
	badger "github.com/dgraph-io/badger/v2"
)

// DB instance of the BadgerDB
var DB *badger.DB
