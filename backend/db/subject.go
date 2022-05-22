package db

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

func GetAllSubjects() []Subject {
	subjects := []Subject{}
	GetDbConnection().Find(&subjects)
	return subjects
}

func (p *Subject) Persist(tx *gorm.DB) {
	tx.Create(p)
}

func (p *Subject) ToDto() dto.Subject {
	return dto.Subject{
		Name: p.Name,
	}
}
