package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func InitDbConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}
