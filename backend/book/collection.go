package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
	"os"
)

type Collection struct {
	gorm.Model
	Title string `gorm:"uniqueIndex;not null"`
	Cover *string
	Books []*Book `gorm:"foreignKey:CollectionId;references:ID"`
}

func (c *Collection) Persist(tx *gorm.DB) {
	tx.Create(c)
}

func (c *Collection) Updates(tx *gorm.DB) {
	tx.Updates(c)
}

func (c *Collection) ToDto() dto.Collection {
	var books = make([]dto.Book, len(c.Books))
	for i, book := range c.Books {
		books[i] = book.ToDto()
	}
	var cover *[]byte
	if c.Cover != nil {
		readCover, _ := os.ReadFile(*c.Cover)
		cover = &readCover
	}
	return dto.Collection{
		ID:    c.ID,
		Cover: cover,
		Title: c.Title,
		Books: books,
	}
}

func GetCollectionByName(name string) Collection {
	var collection Collection
	db.GetDbConnection().Preload("Books").Preload("Books.Authors").Preload("Books.Subjects").Find(&collection, "title = ?", name)
	return collection
}

func GetLazyCollectionByName(name string, tx *gorm.DB) Collection {
	var collection Collection
	tx.Preload("Books").Preload("Books.Authors").Preload("Books.Subjects").Find(&collection, "title = ?", name)
	return collection
}

func GetCollectionById(id uint64) Collection {
	var collection Collection
	db.GetDbConnection().Preload("Books").Find(&collection, id)
	return collection
}
