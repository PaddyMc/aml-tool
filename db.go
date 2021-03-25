package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// DB will contain the read/write db logic
type DB struct {
	leveldb *leveldb.DB
}

// NewDB returns a new instance of the db
func NewDB(path string) *DB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		fmt.Println(err)
	}

	return &DB{
		leveldb: db,
	}
}

// GetData returns data is a given key exists
func (db *DB) GetData(key []byte) (string, error) {
	value, err := db.leveldb.Get(key, nil)

	return string(value), err
}

// PutData puts key, value data into the database
func (db *DB) PutData(key []byte, value []byte) error {
	err := db.leveldb.Put(key, value, nil)

	return err
}
