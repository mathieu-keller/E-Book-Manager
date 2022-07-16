package db

import (
	"e-book-manager/dto"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Book struct {
	gorm.Model
	Title            string `gorm:"index:idx_book_title,unique;not null"`
	Published        *time.Time
	Language         string
	Subjects         []*Subject `gorm:"many2many:Subject2Book;"`
	Publisher        *string
	Cover            *string
	BookPath         string
	OriginalBookName string
	OriginalBookPath string
	Authors          []*Author `gorm:"many2many:Author2Book;"`
	CollectionId     *uint
	CollectionIndex  *uint
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

func (p *Book) Update(tx *gorm.DB) error {
	err := tx.Model(p).
		Association("Subjects").
		Replace(p.Subjects)
	if err != nil {
		return err
	}
	err = tx.Model(p).
		Association("Authors").
		Replace(p.Authors)
	if err != nil {
		return err
	}
	tx.Updates(p)
	return nil
}

func (p *Book) ToDto() dto.Book {
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
		Cover:           p.Cover,
		Book:            p.BookPath,
		CollectionId:    p.CollectionId,
		CollectionIndex: p.CollectionIndex,
		Authors:         authors,
	}
}

func GetBookByTitle(title string) Book {
	var book Book
	GetDbConnection().
		Preload("Authors").
		Preload("Subjects").
		Find(&book, "title = ?", title)
	return book
}

func GetBookById(id string) Book {
	var book Book
	GetDbConnection().
		Find(&book, "id = ?", id)
	return book
}

func SearchBooks(search []string, page int) []Book {
	var books []Book
	var searchQuery = make([]string, len(search))
	for i, s := range search {
		searchQuery[i] = "search_terms ilike '%" + s + "%'"
	}
	GetDbConnection().
		Offset(SetPage(page)).
		Limit(Limit).
		Table("books_search").
		Find(&books, strings.Join(searchQuery, " and "))
	return books
}
