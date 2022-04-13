package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex;not null"`
	Books []*Book `gorm:"many2many:Subject2Book;"`
}

func GetSubjectByName(name string) Subject {
	var subject = Subject{}
	db.GetDbConnection().Find(&subject, "name = ?", name)
	return subject
}

func (p *Subject) Persist() {
	db.GetDbConnection().Create(p)
}

func (p *Subject) ToDto() dto.Subject {
	books := make([]dto.Book, len(p.Books))
	for i, book := range p.Books {
		books[i] = book.ToDto()
	}
	return dto.Subject{
		Name:  p.Name,
		Books: books,
	}
}
