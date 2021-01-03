package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// DB generic db
type DB struct {
	Client *gorm.DB
}

// NewSqliteDB create a sqlite db connection
func NewSqliteDB(dbpath string) *DB {

	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}
