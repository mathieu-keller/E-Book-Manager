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
	savedBook := Book{}
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
	var books = make([]Book, 0)
	var where = ""
	for _, s := range search {
		where += "(books.title LIKE '%" + s + "%' OR " +
			" authors.name LIKE '%" + s + "%' OR " +
			" collections.title LIKE '%" + s + "%' OR " +
			" subjects.name LIKE '%" + s + "%') and "
	}
	db.GetDbConnection().Offset(db.SetPage(page)).Limit(db.Limit).Table("books").Joins("left join collections on books.collection_id = collections.id" +
		"").Joins("left join author2_books on author2_books.book_id = books.id" +
		"").Joins("left join authors on authors.id = author2_books.author_id" +
		"").Joins("left join subject2_books on subject2_books.book_id = books.id" +
		"").Joins("left join subjects on subjects.id = subject2_books.subject_id" +
		"").Where(where +
		" 1=1").Group("books.title" +
		"").Find(&books)
	return books
}
