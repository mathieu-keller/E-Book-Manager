package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB
var Limit = 32

func GetDbConnection() *gorm.DB {
	if database == nil {
		dsn := "host=" + os.Getenv("dbAddress") + " user=" + os.Getenv("dbUser") + " password=" + os.Getenv("dbPassword") + " dbname=" + os.Getenv("dbName") + " port=" + os.Getenv("dbPort") + " sslmode=disable TimeZone=Europe/Berlin"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic("failed to connect database")
		}
		database = db
	}
	return database
}

func Migrate() {
	dbCon := GetDbConnection()
	err := dbCon.AutoMigrate(&Book{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&Author{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&Subject{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&Collection{})
	if err != nil {
		panic(err.Error())
	}
}

func SetPage(page int) int {
	return (page - 1) * Limit
}
