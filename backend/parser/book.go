package parser

import (
	"e-book-manager/db"
	"e-book-manager/epub/epubReader"
	"e-book-manager/epub/epubWriter"
	"errors"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadFile(fileHeader *multipart.FileHeader) error {
	bookFile, err := epubReader.Open("upload/tmp/" + fileHeader.Filename)
	if err != nil {
		return err
	}
	defer bookFile.Close()
	return ParseBook(bookFile, "upload/tmp/", fileHeader.Filename)
}

func ParseBook(epubBook *epubReader.Book, originalFilePath string, originalFileName string) error {
	if epubBook.Opf.Metadata == nil {
		return errors.New("no metadata found")
	}
	tx := db.GetDbConnection().Begin()
	metadata := *epubBook.Opf.Metadata
	var coverId = ""
	var metaIdMap = make(map[string]map[string]epubReader.Meta)
	if metadata.Meta != nil {
		for _, meta := range *metadata.Meta {
			if meta.Name == "cover" {
				coverId = meta.Content
			} else if meta.Refines != "" {
				if metaIdMap[meta.Refines] == nil {
					metaIdMap[meta.Refines] = make(map[string]epubReader.Meta)
				}
				metaIdMap[meta.Refines][meta.Property] = meta
			}
		}
	}
	bookEntity := db.Book{}
	bookEntity.Title = GetTitle(metadata, metaIdMap)
	if bookEntity.Title == "" {
		tx.Rollback()
		return errors.New("no title found")
	}
	bookEntity.Authors = GetAuthor(metadata, metaIdMap, tx)
	bookEntity.Published, _ = GetDate(metadata)
	bookEntity.Publisher, _ = GetPublisher(metadata)
	bookEntity.Language, _ = GetLanguage(metadata)
	bookEntity.Subjects = GetSubject(metadata, tx)
	bookEntity.CollectionIndex = GetCollectionIndex(metadata)
	bookEntity.Persist(tx)
	filePath := "upload/ebooks/" + strconv.Itoa(int(bookEntity.ID)) + "-" + bookEntity.Title + "/"
	err := os.MkdirAll(filePath, 0770)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = epubWriter.CopyZip(epubBook, filePath)
	if err != nil {
		tx.Rollback()
		return err
	}
	bookEntity.BookPath = filePath + "book.epub"
	bookEntity.OriginalBookPath = filePath + "original.epub"
	bookEntity.OriginalBookName = originalFileName
	bookEntity.Cover, _ = GetCover(coverId, epubBook, filePath)
	bookEntity.CollectionId = GetCollection(metadata, metaIdMap, bookEntity.Cover, tx)
	bookEntity.Update(tx)
	err = os.Rename(originalFilePath+originalFileName, filePath+"original.epub")
	if err != nil {
		tx.Rollback()
		return err
	}
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	return nil
}
