package book

import (
	"e-book-manager/db"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string  `gorm:"unique;index"`
	Books []*Book `gorm:"many2many:Author2Book;"`
}

func GetAuthorByName(name string) Author {
	var author = Author{}
	db.GetDbConnection().Find(&author, "name = ?", name)
	return author
}
