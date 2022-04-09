package book

import (
	"e-book-manager/db"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Name         string `gorm:"unique;index"`
	Published    time.Time
	Language     string
	Subject      string
	Publisher    string
	Cover        string
	Book         string
	Author       []*Author `gorm:"many2many:Author2Book;"`
	CollectionId uint
}

func (p *Book) Persist() {
	db.GetDbConnection().Create(p)
}

func GetAllBooks() []Book {
	var books []Book
	db.GetDbConnection().Find(&books, "")
	return books
}
