package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
	"os"
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

func (p *Book) ToDto() dto.Book {
	cover, _ := os.ReadFile(p.Cover)
	return dto.Book{
		ID:           p.ID,
		Name:         p.Name,
		Published:    p.Published,
		Language:     p.Language,
		Subject:      p.Subject,
		Publisher:    p.Publisher,
		Cover:        cover,
		Book:         p.Book,
		CollectionId: p.CollectionId,
	}
}

func GetBookByTitle(title string) Book {
	var book Book
	db.GetDbConnection().Find(&book, "name = ?", title)
	return book
}

func GetBookById(id string) Book {
	var book Book
	db.GetDbConnection().Find(&book, "ID = ?", id)
	return book
}
