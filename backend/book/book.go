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
	Title        string `gorm:"index:idx_book_title,unique;not null"`
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
	var savedBook Book
	connection := db.GetDbConnection()
	connection.Find(&savedBook, "title = ?", p.Title)
	if savedBook.ID == 0 {
		connection.Create(p)
	} else {
		p.ID = savedBook.ID
		connection.Updates(p)
	}
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
	db.GetDbConnection().Preload("Authors").Preload("Subjects").Find(&book, "title = ?", title)
	return book
}

func GetBookById(id string) Book {
	var book Book
	db.GetDbConnection().Find(&book, "id = ?", id)
	return book
}

func SearchBooks(search []string, page int) []Book {
	var books []Book
	selector := db.GetDbConnection().Offset(db.SetPage(page)).Limit(db.Limit).Table("books" +
		"").Joins("left join collections on books.collection_id = collections.id")
	for _, s := range search {
		selector.Where("books.title ILIKE ? OR "+
			" collections.title ILIKE ? or "+
			" (SELECT count(*) from subjects s JOIN subject2_books S2B ON S.ID = S2B.SUBJECT_ID "+
			" WHERE s2b.BOOK_ID = books.ID and s.NAME ILIKE ?) >= 1 or "+
			" (SELECT count(*) from authors a JOIN author2_books A2B ON a.ID = A2B.author_id "+
			" WHERE A2B.BOOK_ID = books.ID and a.NAME ILIKE ?) >= 1",
			"%"+s+"%", "%"+s+"%", "%"+s+"%", "%"+s+"%")
	}
	selector.Distinct().Find(&books)
	return books
}
