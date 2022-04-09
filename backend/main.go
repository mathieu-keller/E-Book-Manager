package main

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/parser/epub"
	"errors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func uploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("myFile")
	if err != nil {
		c.String(400, err.Error())
		return
	}
	var dataType = fileHeader.Header.Values("Content-Type")[0]
	if dataType != "application/epub+zip" {
		c.String(400, "wrong data type: "+dataType)
		return
	}
	os.MkdirAll("upload/ebooks/", os.ModePerm)
	err = c.SaveUploadedFile(fileHeader, "upload/ebooks/"+fileHeader.Filename)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	bookFile, err := epub.Open("upload/ebooks/" + fileHeader.Filename)
	defer bookFile.Close()

	err = createBookEntity(bookFile)

	c.String(200, "OK!")
}

// todo error handling?
func createBookEntity(bookFile *epub.Book) error {
	var coverId = ""
	var metaIdMap = make(map[string]map[string]epub.Metafield)
	for _, meta := range bookFile.Opf.Metadata.Meta {
		if meta.Name == "cover" {
			coverId = meta.Content
		} else if meta.Refines != "" {
			if metaIdMap[meta.Refines] == nil {
				metaIdMap[meta.Refines] = make(map[string]epub.Metafield)
			}
			metaIdMap[meta.Refines][meta.Property] = meta
		}
	}
	bookEntity := book.Book{}
	setAuthor(bookFile, &bookEntity, metaIdMap)
	var date, _ = getDate(bookFile)
	bookEntity.Published = *date
	bookEntity.Publisher, _ = getPublisher(bookFile)
	bookEntity.Language, _ = getLanguage(bookFile)
	setTitles(bookFile, metaIdMap, &bookEntity)

	bookEntity.Cover, _ = getCover(coverId, bookFile, bookEntity.Name)
	bookEntity.Persist()
	return nil
}

func setAuthor(bookFile *epub.Book, bookEntity *book.Book, metaIdMap map[string]map[string]epub.Metafield) {
	for _, creator := range bookFile.Opf.Metadata.Creator {
		var ele = metaIdMap["#"+creator.ID]
		if ele["role"].Data == "aut" {
			var authorName = strings.TrimSpace(creator.Data)
			var author = book.GetAuthorByName(authorName)
			if author.Name == "" {
				author.Name = authorName
			}
			bookEntity.Author = append(bookEntity.Author, &author)
		}
	}
}

func getCover(coverId string, bookFile *epub.Book, bookName string) (string, error) {
	var href = ""
	var imgTyp = ""
	if coverId != "" {
		for _, mani := range bookFile.Opf.Manifest {
			if mani.ID == coverId {
				href = mani.Href
				if mani.MediaType == "image/gif" {
					imgTyp = ".gif"
				} else if mani.MediaType == "image/jpeg" {
					imgTyp = ".jpg"
				} else if mani.MediaType == "image/png" {
					imgTyp = ".png"
				} else if mani.MediaType == "image/svg+xml" {
					imgTyp = ".svg"
				}
				break
			}
		}
	}
	if href != "" {
		readedFile, err := bookFile.Open(href)
		if err != nil {
			return "", err
		}
		defer readedFile.Close()
		b, err := ioutil.ReadAll(readedFile)
		if err != nil {
			return "", err
		}
		var path = "upload/covers/" + bookName + "/"
		os.MkdirAll(path, os.ModePerm)
		err = ioutil.WriteFile(path+"cover"+imgTyp, b, fs.ModePerm)
		if err != nil {
			return "", err
		}
		return path + "cover" + imgTyp, nil
	}
	return "", errors.New("cover not found!")
}

func setTitles(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Metafield, bookEntity *book.Book) {
	for _, titleMeta := range bookFile.Opf.Metadata.Title {
		if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "main" {
			bookEntity.Name = titleMeta.Data
		} else if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "collection" {
			var collectionName = strings.TrimSpace(titleMeta.Data)
			var collection = book.GetCollectionByName(collectionName)
			if collection.Name == "" {
				collection.Name = collectionName
				collection.Persist()
			}
			bookEntity.CollectionId = collection.ID
		}
	}
}

func getLanguage(bookFile *epub.Book) (string, error) {
	lang := bookFile.Opf.Metadata.Language
	if len(lang) != 1 {
		return "", errors.New("multi lang not supported!")
	}
	return lang[0], nil
}

func getPublisher(bookFile *epub.Book) (string, error) {
	pub := bookFile.Opf.Metadata.Publisher
	if len(pub) != 1 {
		return "", errors.New("More then 1 publisher found")
	}
	return pub[0], nil
}

func getDate(bookFile *epub.Book) (*time.Time, error) {
	date := bookFile.Opf.Metadata.Date
	if len(date) != 1 {
		return nil, errors.New("multi date not supported!")
	}
	var time, err = time.Parse("2006-01-02", date[0].Data)
	return &time, err
}

func setupRoutes() {
	r := gin.Default()
	r.POST("/upload", uploadFile)
	r.GET("/", func(c *gin.Context) {
		file, err := os.ReadFile("index.html")
		if err != nil {
			c.String(500, err.Error())
		}
		c.Data(200, "text/html; charset=utf-8", file)
	})
	r.GET("/all", func(c *gin.Context) {
		c.JSON(200, book.GetAllBooks())
	})
	r.Run()
}

func main() {
	//todo maybe in a different place?
	db := db.GetDbConnection()
	// Migrate the schema
	db.AutoMigrate(&book.Book{})
	db.AutoMigrate(&book.Author{})
	db.AutoMigrate(&book.Collection{})
	setupRoutes()
}
