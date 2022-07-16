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
	GetDbConnection().
		Offset(SetPage(page)).
		Limit(Limit).
		Table("library_items").
		Find(&libraryItems)
	return libraryItems
}
