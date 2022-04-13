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
	Title        string `gorm:"uniqueIndex;not null"`
	Published    time.Time
	Language     string
	Subjects     []*Subject `gorm:"many2many:Subject2Book;"`
	Publisher    string
	Cover        string
	Book         string
	Authors      []*Author `gorm:"many2many:Author2Book;"`
	CollectionId uint
}

func (p *Book) Persist() {
	db.GetDbConnection().Create(p)
}

func (p *Book) ToDto() dto.Book {
	cover, _ := os.ReadFile(p.Cover)
	subjects := make([]dto.Subject, len(p.Subjects))
	for i, subject := range p.Subjects {
		subjects[i] = subject.ToDto()
	}
	authors := make([]dto.Author, len(p.Authors))
	for i, author := range p.Authors {
		authors[i] = author.ToDto()
	}
	return dto.Book{
		ID:           p.ID,
		Title:        p.Title,
		Published:    p.Published,
		Language:     p.Language,
		Subjects:     subjects,
		Publisher:    p.Publisher,
		Cover:        cover,
		Book:         p.Book,
		CollectionId: p.CollectionId,
		Authors:      authors,
	}
}

func GetBookByTitle(title string) Book {
	var book Book
	db.GetDbConnection().Find(&book, "title = ?", title)
	return book
}

func GetBookById(id string) Book {
	var book Book
	db.GetDbConnection().Find(&book, "ID = ?", id)
	return book
}
