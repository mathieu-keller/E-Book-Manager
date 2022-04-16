package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Title string  `gorm:"uniqueIndex;not null"`
	Books []*Book `gorm:"foreignKey:CollectionId;references:ID"`
}

func (c *Collection) Persist() {
	db.GetDbConnection().Create(c)
}

func (c *Collection) ToDto() dto.Collection {
	var books = make([]dto.Book, len(c.Books))
	for i, book := range c.Books {
		books[i] = book.ToDto()
	}
	return dto.Collection{
		ID:    c.ID,
		Title: c.Title,
		Books: books,
	}
}

func GetCollectionByName(name string) Collection {
	var collection = Collection{}
	db.GetDbConnection().Preload("books").Preload("books.authors").Preload("books.subjects").Find(&collection, "title = ?", name)
	return collection
}

func GetCollectionById(id uint64) Collection {
	var collection = Collection{}
	db.GetDbConnection().Preload("books").Find(&collection, id)
	return collection
}
