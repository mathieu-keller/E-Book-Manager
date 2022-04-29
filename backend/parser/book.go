package parser

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/epub"
	"e-book-manager/epub/convert"
	"errors"
	"strconv"
)

func ParseBook(epubBook *epub.Book) error {
	if epubBook.Opf.Metadata == nil {
		return errors.New("no metadata found")
	}
	tx := db.GetDbConnection().Begin()
	metadata := *epubBook.Opf.Metadata
	var coverId = ""
	var metaIdMap = make(map[string]map[string]epub.Meta)
	if metadata.Meta != nil {
		for _, meta := range *metadata.Meta {
			if meta.Name == "cover" {
				coverId = meta.Content
			} else if meta.Refines != "" {
				if metaIdMap[meta.Refines] == nil {
					metaIdMap[meta.Refines] = make(map[string]epub.Meta)
				}
				metaIdMap[meta.Refines][meta.Property] = meta
			}
		}
	}
	bookEntity := book.Book{}
	bookEntity.Title = GetTitle(metadata, metaIdMap)
	if bookEntity.Title == "" {
		tx.Rollback()
		return errors.New("no title found")
	}
	bookEntity.Authors = GetAuthor(metadata, metaIdMap, tx)
	bookEntity.Published, _ = GetDate(metadata)
	bookEntity.Publisher, _ = GetPublisher(metadata)
	bookEntity.Language, _ = GetLanguage(metadata)
	bookEntity.Cover, _ = GetCover(coverId, epubBook, bookEntity.Title)
	bookEntity.Subjects = GetSubject(metadata, tx)
	bookEntity.CollectionIndex = GetCollectionIndex(metadata)
	bookEntity.CollectionId = GetCollection(metadata, metaIdMap, bookEntity.Cover, tx)
	bookEntity.Persist(tx)
	filePath := "upload/ebooks/" + strconv.Itoa(int(bookEntity.ID)) + "-" + bookEntity.Title + ".epub"
	bookEntity.Book = filePath
	bookEntity.Update(tx)

	err := convert.CopyZip(epubBook, filePath)
	if tx.Error != nil || err != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil
}
