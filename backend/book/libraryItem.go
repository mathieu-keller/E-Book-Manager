package book

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"os"
)

type LibraryItem struct {
	Cover     string
	Title     string
	ItemType  string
	BookCount uint
	Id        uint
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	cover, _ := os.ReadFile(p.Cover)
	return dto.LibraryItem{
		Cover:     cover,
		Title:     p.Title,
		ItemType:  p.ItemType,
		BookCount: p.BookCount,
		Id:        p.Id,
	}
}

func GetAllLibraryItems() []LibraryItem {
	var libraryItems = make([]LibraryItem, 0)
	db.GetDbConnection().Limit(32).Table("books").Select(" COALESCE(collections.id, books.id) as id, " +
		" books.cover as cover, COALESCE(collections.title, books.title) AS title, " +
		" CASE WHEN collections.title IS NOT NULL THEN 'collection' " +
		" ELSE 'book' END AS itemType, COUNT(*) AS bookCount " +
		"").Joins("left join collections on books.collection_id = collections.id" +
		"").Group("COALESCE(collections.title, books.title)" +
		"").Scan(&libraryItems)
	return libraryItems
}

func GetLibraryItemByCollectionId(id uint64) LibraryItem {
	var libraryItem = LibraryItem{}
	db.GetDbConnection().Table("BOOKS").Select(" COALESCE(collections.id, books.id) as id, "+
		" books.cover as cover, COALESCE(collections.title, books.title) AS title, "+
		" CASE WHEN collections.title IS NOT NULL THEN 'collection' "+
		" ELSE 'book' END AS itemType, COUNT(*) AS bookCount "+
		"").Joins("left join collections on books.collection_id = collections.id"+
		"").Group("COALESCE(collections.title, books.title)"+
		"").Find(&libraryItem, "collections.id = ?", id)
	return libraryItem
}
