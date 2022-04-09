package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"os"
)

type LibraryItem struct {
	ID    uint
	Name  string
	Type  string
	Cover string
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	cover, _ := os.ReadFile(p.Cover)
	return dto.LibraryItem{
		ID:    p.ID,
		Name:  p.Name,
		Cover: cover,
	}
}

func GetAllLibraryItems() []LibraryItem {
	var libraryItems = make([]LibraryItem, 15)
	db.GetDbConnection().Table("BOOKS").Select("BOOKS.Cover as Cover, " +
		"COALESCE(collections.ID, BOOKS.ID) AS ID, " +
		"COALESCE(collections.NAME, BOOKS.NAME) AS Name, " +
		"CASE WHEN collections.NAME IS NULL THEN 'BOOK' ELSE 'COLLECTION' END AS Type").Joins("left join " +
		"collections on BOOKS.COLLECTION_ID = collections.id").Group("COALESCE(collections.NAME, BOOKS.NAME)").Scan(&libraryItems)
	return libraryItems
}
