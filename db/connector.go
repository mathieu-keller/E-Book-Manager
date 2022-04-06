package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
