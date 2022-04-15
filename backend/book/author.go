package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex;not null"`
	Books []*Book `gorm:"many2many:Author2Book;"`
}

func GetAuthorByName(name string) Author {
	var author = Author{}
	db.GetDbConnection().Find(&author, "name = ?", name)
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
