package main

import (
	"e-book-manager/resources"
	"os"
)

func main() {
	err := os.MkdirAll("upload/ebooks/", 0770)
	if err != nil {
		panic(err.Error())
	}
	resources.SetupRoutes()
}
