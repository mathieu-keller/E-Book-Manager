package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"os"
)

type LibraryItem struct {
	Cover     string
	Name      string
	ItemType  string
	BookCount uint
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	cover, _ := os.ReadFile(p.Cover)
	return dto.LibraryItem{
		Cover:     cover,
		Name:      p.Name,
		ItemType:  p.ItemType,
		BookCount: p.BookCount,
	}
}

func GetAllLibraryItems() []LibraryItem {
	var libraryItems = make([]LibraryItem, 0)
	db.GetDbConnection().Table("BOOKS").Select("BOOKS.COVER as Cover, COALESCE(COLLECTIONS.NAME, BOOKS.NAME) AS Name, " +
		" CASE WHEN COLLECTIONS.NAME IS NOT NULL THEN 'collection' " +
		" ELSE 'book' END AS ItemType, COUNT(*) AS BookCount " +
		"").Joins("left join COLLECTIONS on BOOKS.COLLECTION_ID = COLLECTIONS.id" +
		"").Group("COALESCE(collections.NAME, BOOKS.NAME)" +
		"").Scan(&libraryItems)
	return libraryItems
}
