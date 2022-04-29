package book

import (
	"e-book-manager/dto"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex;not null"`
	Books []*Book `gorm:"many2many:Author2Book;"`
}

func (a *Author) Create(tx *gorm.DB) {
	tx.Create(a)
}

func GetAuthorByName(name string, tx *gorm.DB) Author {
	var author = Author{}
	tx.Find(&author, "name = ?", name)
	return author
}

func (p *Author) ToDto() dto.Author {
	books := make([]dto.Book, len(p.Books))
	for i, book := range p.Books {
		books[i] = book.ToDto()
	}
	return dto.Author{
		ID:    p.ID,
		Name:  p.Name,
		Books: books,
	}
}
