package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var database *gorm.DB
var Limit = 32

func GetDbConnection() *gorm.DB {
	if database == nil {
		dsn := "host=" + os.Getenv("dbAddress") + " user=" + os.Getenv("dbUser") + " password=" + os.Getenv("dbPassword") + " dbname=" + os.Getenv("dbName") + " port=" + os.Getenv("dbPort") + " sslmode=disable TimeZone=Europe/Berlin"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
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
	GetDbConnection().Exec("drop view library_items")
	GetDbConnection().Exec("drop view books_search")

	err := dbCon.AutoMigrate(&Book{}, &Author{}, &Subject{}, &Collection{})
	if err != nil {
		panic(err.Error())
	}
	query, err := ioutil.ReadFile("sql/library_items_view.sql")
	if err != nil {
		panic(err.Error())
	}
	GetDbConnection().Exec(string(query))
	query, err = ioutil.ReadFile("sql/books_search_view.sql")
	if err != nil {
		panic(err.Error())
	}
	GetDbConnection().Exec(string(query))

}

func SetPage(page int) int {
	return (page - 1) * Limit
}
