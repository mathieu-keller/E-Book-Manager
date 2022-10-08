package parser

import (
	"archive/zip"
	"e-book-manager/db"
	"e-book-manager/dto"
	"e-book-manager/epub/epubReader"
	"e-book-manager/epub/epubWriter"
	"errors"
	"gorm.io/gorm"
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
	tx := db.GetDbConnection().Begin()
	if epubBook.Opf.Metadata == nil {
		return errors.New("no metadata found")
	}
	metadata, metaIdMap, coverId := getMetadata(epubBook)
	bookEntity := &db.Book{}
	err := fillBookEntity(bookEntity, metadata, metaIdMap, epubBook, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	bookEntity.Persist(tx)
	filePath := "upload/ebooks/" + strconv.Itoa(int(bookEntity.ID)) + "-" + bookEntity.Title + "/"
	err = os.MkdirAll(filePath, 0770)
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
	bookEntity.Cover, _ = GetCover(coverId, epubBook)
	bookEntity.CollectionId = GetCollection(metadata, metaIdMap, bookEntity.Cover, tx)
	err = bookEntity.Update(tx)
	if err != nil {
		return err
	}
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

func getMetadata(epubBook *epubReader.Book) (epubReader.Metadata, map[string]map[string]epubReader.Meta, string) {
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
	return metadata, metaIdMap, coverId
}

func fillBookEntity(bookEntity *db.Book, metadata epubReader.Metadata, metaIdMap map[string]map[string]epubReader.Meta, book *epubReader.Book, tx *gorm.DB) error {
	title, err := GetTitle(metadata)
	if err != nil {
		return err
	}
	bookEntity.Title = title
	creator, err := GetAuthor(metadata, tx)
	if err != nil {
		return err
	}
	authors := make([]*db.Author, 1)
	authors = append(authors, creator)
	bookEntity.Authors = authors
	bookEntity.Published, _ = GetDate(metadata)
	bookEntity.Publisher, _ = GetPublisher(metadata)
	bookEntity.Language, _ = GetLanguage(metadata)
	bookEntity.Subjects = GetSubject(metadata, tx)
	bookEntity.CollectionIndex = GetCollectionIndex(metadata)
	return nil
}

func UpdateBookData(bookDto dto.Book, tx *gorm.DB) error {
	book := db.GetBookByTitle(bookDto.Title)
	zipReader, err := zip.OpenReader(book.BookPath)
	if err != nil {
		return err
	}
	epub, err := epubReader.Open(&zipReader.Reader)
	if err != nil {
		return err
	}
	subjectNames := make([]epubReader.Subject, len(bookDto.Subjects))
	for i, subject := range bookDto.Subjects {
		subjectNames[i] = epubReader.Subject{Text: subject.Name}
	}
	epub.Opf.Metadata.Subject = &subjectNames

	metadata, metaIdMap, _ := getMetadata(epub)
	err = fillBookEntity(&book, metadata, metaIdMap, epub, tx)
	if err != nil {
		return err
	}
	err = book.Update(tx)
	if err != nil {
		return err
	}
	err = epubWriter.CreateZip(epub, book.BookPath+"copy")
	if err != nil {
		return err
	}
	err = os.Rename(book.BookPath, book.BookPath+"backup")
	if err != nil {
		return err
	}
	err = os.Rename(book.BookPath+"copy", book.BookPath)
	if err != nil {
		os.Rename(book.BookPath+"backup", book.BookPath)
		os.Remove(book.BookPath + "copy")
		return err
	}
	return os.Remove(book.BookPath + "backup")
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
