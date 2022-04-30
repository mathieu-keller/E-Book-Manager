package main

import (
	"e-book-manager/db"
	"e-book-manager/resources"
	"os"
)

func main() {
	err := os.MkdirAll("upload/tmp/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = os.MkdirAll("upload/ebooks/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	db.Migrate()
	resources.SetupRoutes()
}
