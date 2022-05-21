package parser

import (
	"archive/zip"
	"e-book-manager/db"
	"e-book-manager/epub/epubReader"
	"e-book-manager/epub/epubWriter"
	"errors"
	"os"
	"strconv"
)

func UploadFile(reader *zip.Reader, orgFileName string) error {
	bookFile, err := epubReader.Open(reader)
	if err != nil {
		return err
	}
	return ParseBook(bookFile, orgFileName)
}

func ParseBook(epubBook *epubReader.Book, originalFileName string) error {
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
	err = epubWriter.CreateZip(epubBook, filePath+"book.epub")
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
	err = saveOriginalBook(epubBook, err, filePath)
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

func saveOriginalBook(epubBook *epubReader.Book, err error, filePath string) error {
	file, err := os.Create(filePath + "original.epub")
	if err != nil {
		return err
	}
	writer := zip.NewWriter(file)
	defer writer.Close()
	for _, file := range epubBook.Fd.File {
		if file.FileInfo().IsDir() {
			continue
		}
		err := writer.Copy(file)
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
