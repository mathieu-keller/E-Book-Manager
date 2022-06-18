package db

import (
	"e-book-manager/dto"
)

type LibraryItem struct {
	Cover     *string
	Title     string
	ItemType  string
	BookCount uint
	Id        uint
}

func (p *LibraryItem) ToDto() dto.LibraryItem {
	return dto.LibraryItem{
		Cover:     p.Cover,
		Title:     p.Title,
		ItemType:  p.ItemType,
		BookCount: p.BookCount,
		Id:        p.Id,
	}
}

func GetAllLibraryItems(page int) []LibraryItem {
	var libraryItems []LibraryItem
	GetDbConnection().Offset(SetPage(page)).Limit(Limit).Table("books" +
		"").Select("" +
		" COALESCE(collections.id, books.id) as Id, " +
		" COALESCE(collections.title, books.title) as Title, " +
		" CASE WHEN collections.title IS NOT NULL THEN 'collection' ELSE 'book' END AS Item_Type, " +
		" COUNT(*) AS Book_Count, " +
		" CASE WHEN collections.title is not null then collections.Cover ELSE books.Cover END as Cover" +
		"").Joins("left join collections on books.collection_id = collections.id" +
		"").Group("COALESCE(collections.id, books.id), COALESCE(collections.title, books.title), " +
		" collections.title, CASE WHEN collections.title is not null then collections.Cover ELSE books.Cover END" +
		"").Scan(&libraryItems)
	return libraryItems
}
