package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type Book struct {
	gorm.Model
	Title           string `gorm:"index:idx_book_title,unique;not null"`
	Published       time.Time
	Language        string
	Subjects        []*Subject `gorm:"many2many:Subject2Book;"`
	Publisher       *string
	Cover           *string
	Book            string
	Authors         []*Author `gorm:"many2many:Author2Book;"`
	CollectionId    *uint
	CollectionIndex *uint
}

func (p *Book) Persist(tx *gorm.DB) {
	var savedBook Book
	tx.Find(&savedBook, "title = ?", p.Title)
	if savedBook.ID == 0 {
		tx.Create(p)
	} else {
		p.ID = savedBook.ID
		tx.Updates(p)
	}
}

func (p *Book) Update(tx *gorm.DB) {
	tx.Updates(p)
}

func (p *Book) ToDto() dto.Book {
	var cover *[]byte
	if p.Cover != nil {
		readCover, _ := os.ReadFile(*p.Cover)
		cover = &readCover
	}
	subjects := make([]dto.Subject, len(p.Subjects))
	for i, subject := range p.Subjects {
		subjects[i] = subject.ToDto()
	}
	authors := make([]dto.Author, len(p.Authors))
	for i, author := range p.Authors {
		authors[i] = author.ToDto()
	}
	return dto.Book{
		ID:              p.ID,
		Title:           p.Title,
		Published:       p.Published,
		Language:        p.Language,
		Subjects:        subjects,
		Publisher:       p.Publisher,
		Cover:           cover,
		Book:            p.Book,
		CollectionId:    p.CollectionId,
		CollectionIndex: p.CollectionIndex,
		Authors:         authors,
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
		"").Joins(" left join collections on books.collection_id = collections.id ")
	for i, s := range search {
		stringIndex := strconv.Itoa(i)
		selector.Joins(" left JOIN SUBJECT2_BOOKS  AS s2b" + stringIndex + " ON books.ID = s2b" + stringIndex + ".BOOK_ID " +
			" left JOIN SUBJECTS  AS S" + stringIndex + " ON S" + stringIndex + ".ID = s2b" + stringIndex + ".SUBJECT_ID" +
			" left JOIN AUTHOR2_BOOKS AS a2b" + stringIndex + " ON books.ID = a2b" + stringIndex + ".BOOK_ID " +
			" left JOIN AUTHORS AS a" + stringIndex + " ON a" + stringIndex + ".ID = a2b" + stringIndex + ".AUTHOR_ID ")
		selector.Where("books.title ILIKE ? OR "+
			" collections.title ILIKE ? or "+
			" S"+stringIndex+".name ILIKE ? or "+
			" a"+stringIndex+".name ILIKE ? ",
			"%"+s+"%", "%"+s+"%", "%"+s+"%", "%"+s+"%")
	}
	selector.Distinct().Find(&books)
	return books
}
