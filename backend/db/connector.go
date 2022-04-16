package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func GetDbConnection() *gorm.DB {
	if database == nil {
		dsn := os.Getenv("dbUser") + ":" + os.Getenv("dbPassword") +
			"@tcp(" + os.Getenv("dbAddress") +
			")/" + os.Getenv("dbName") + "?charset=utf8mb4&parseTime=True&loc=Local"
		mariaDb := mysql.Open(dsn)
		db, err := gorm.Open(mariaDb, &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		database = db
	}
	return database
}
