package book

import (
	"e-book-manager/dto"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name  string  `gorm:"uniqueIndex;not null"`
	Books []*Book `gorm:"many2many:Subject2Book;"`
}

func GetSubjectByName(name string, tx *gorm.DB) Subject {
	var subject Subject
	tx.Find(&subject, "name = ?", name)
	return subject
}

func (p *Subject) Persist(tx *gorm.DB) {
	tx.Create(p)
}

func (p *Subject) ToDto() dto.Subject {
	books := make([]dto.Book, len(p.Books))
	for i, book := range p.Books {
		books[i] = book.ToDto()
	}
	return dto.Subject{
		ID:    p.ID,
		Name:  p.Name,
		Books: books,
	}
}
