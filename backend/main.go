package main

import (
	"e-book-manager/db"
	"e-book-manager/resources"
	"os"
)

func main() {
	err := os.MkdirAll("upload/tmp/", 0770)
	if err != nil {
		panic(err.Error())
	}
	err = os.MkdirAll("upload/ebooks/", 0770)
	if err != nil {
		panic(err.Error())
	}
	db.Migrate()
	resources.SetupRoutes()
}
