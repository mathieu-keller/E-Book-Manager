package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDbConnection() *gorm.DB {
	dsn := os.Getenv("dbUser") + ":" + os.Getenv("dbPassword") +
		"@tcp(" + os.Getenv("dbAddress") +
		")/" + os.Getenv("dbName") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
