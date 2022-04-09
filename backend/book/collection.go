package book

import (
	"e-book-manager/db"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name  string  `gorm:"unique;index"`
	Books []*Book `gorm:"foreignKey:CollectionId;references:ID"`
}

func (p *Collection) Persist() {
	db.GetDbConnection().Create(p)
}

func GetCollectionByName(name string) Collection {
	var collection = Collection{}
	db.GetDbConnection().Find(&collection, "name = ?", name)
	return collection
}

func GetAllCollections() []Collection {
	var collections []Collection
	db.GetDbConnection().Find(&collections, "")
	return collections
}
