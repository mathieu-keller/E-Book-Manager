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
	Id        uint
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	cover, _ := os.ReadFile(p.Cover)
	return dto.LibraryItem{
		Cover:     cover,
		Name:      p.Name,
		ItemType:  p.ItemType,
		BookCount: p.BookCount,
		Id:        p.Id,
	}
}

func GetAllLibraryItems() []LibraryItem {
	var libraryItems = make([]LibraryItem, 0)
	db.GetDbConnection().Table("BOOKS").Select(" COALESCE(COLLECTIONS.ID, BOOKS.ID) as Id, " +
		" BOOKS.COVER as Cover, COALESCE(COLLECTIONS.NAME, BOOKS.TITLE) AS Name, " +
		" CASE WHEN COLLECTIONS.NAME IS NOT NULL THEN 'collection' " +
		" ELSE 'book' END AS ItemType, COUNT(*) AS BookCount " +
		"").Joins("left join COLLECTIONS on BOOKS.COLLECTION_ID = COLLECTIONS.id" +
		"").Group("COALESCE(collections.NAME, BOOKS.TITLE)" +
		"").Scan(&libraryItems)
	return libraryItems
}

func GetLibraryItemByCollectionId(id uint64) LibraryItem {
	var libraryItem = LibraryItem{}
	db.GetDbConnection().Table("BOOKS").Select(" COALESCE(COLLECTIONS.ID, BOOKS.ID) as Id, "+
		" BOOKS.COVER as Cover, COALESCE(COLLECTIONS.NAME, BOOKS.TITLE) AS Name, "+
		" CASE WHEN COLLECTIONS.NAME IS NOT NULL THEN 'collection' "+
		" ELSE 'book' END AS ItemType, COUNT(*) AS BookCount "+
		"").Joins("left join COLLECTIONS on BOOKS.COLLECTION_ID = COLLECTIONS.id"+
		"").Group("COALESCE(collections.NAME, BOOKS.TITLE)"+
		"").Find(&libraryItem, "collections.ID = ?", id)
	return libraryItem
}
