package main

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/dto"
	epub2 "e-book-manager/epub"
	"e-book-manager/parser"
	"errors"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func uploadFile(fileHeader *multipart.FileHeader) (*book.Book, error) {
	bookFile, err := epub2.Open("upload/ebooks/" + fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	defer bookFile.Close()
	return createBookEntity(bookFile, "upload/ebooks/"+fileHeader.Filename)
}

// todo error handling?
func createBookEntity(bookFile *epub2.Book, path string) (*book.Book, error) {
	var coverId = ""
	var metaIdMap = make(map[string]map[string]epub2.Metafield)
	for _, meta := range bookFile.Opf.Metadata.Meta {
		if meta.Name == "cover" {
			coverId = meta.Content
		} else if meta.Refines != "" {
			if metaIdMap[meta.Refines] == nil {
				metaIdMap[meta.Refines] = make(map[string]epub2.Metafield)
			}
			metaIdMap[meta.Refines][meta.Property] = meta
		}
	}
	bookEntity := book.Book{}
	bookEntity.Title = parser.GetTitle(bookFile, metaIdMap)
	if bookEntity.Title == "" {
		return nil, errors.New("no title found")
	}
	bookEntity.Authors = parser.GetAuthor(bookFile, metaIdMap)
	var date, err = parser.GetDate(bookFile)
	if err == nil {
		bookEntity.Published = *date
	}
	bookEntity.Publisher, _ = parser.GetPublisher(bookFile)
	bookEntity.Language, _ = parser.GetLanguage(bookFile)
	bookEntity.CollectionId = parser.GetCollection(bookFile, metaIdMap)
	bookEntity.Cover, _ = parser.GetCover(coverId, bookFile, bookEntity.Title)
	bookEntity.Subjects = parser.GetSubject(bookFile)
	bookEntity.Book = path
	bookEntity.Persist()
	return &bookEntity, nil
}

func setupRoutes() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.Use(static.Serve("/", static.LocalFile("./bundles", true)))
	r.POST("/upload/multi", func(c *gin.Context) {
		files, _ := c.MultipartForm()

		for _, fileHeader := range files.File["myFiles"] {

			c.SaveUploadedFile(fileHeader, "upload/ebooks/"+fileHeader.Filename)
			uploadFile(fileHeader)
		}
		c.String(200, "Done")
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("myFile")
		err := c.SaveUploadedFile(file, "upload/ebooks/"+file.Filename)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		entity, err := uploadFile(file)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, entity.ToDto())
	})
	r.GET("/library/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		entity := book.GetLibraryItemByCollectionId(id)
		c.JSON(200, entity.ToDto())
	})
	r.GET("/book", func(c *gin.Context) {
		queryParam, exist := c.GetQuery("q")
		if !exist {
			c.String(400, "query param q expected")
		}
		search := strings.Split(queryParam, " ")
		var books = book.SearchBooks(search)
		bookDtos := make([]dto.Book, len(books))
		for i, b := range books {
			bookDtos[i] = b.ToDto()
		}
		c.JSON(200, bookDtos)
	})
	r.GET("/book/:title", func(c *gin.Context) {
		title := c.Param("title")
		entity := book.GetBookByTitle(title)
		c.JSON(200, entity.ToDto())
	})
	r.GET("/collection", func(c *gin.Context) {
		title := c.Query("title")
		byName := book.GetCollectionByName(title)
		c.JSON(200, byName.ToDto())
	})
	r.GET("/collection/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		byName := book.GetCollectionById(id)
		c.JSON(200, byName.ToDto())
	})
	r.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookEntity := book.GetBookById(id)
		b, err := os.ReadFile(bookEntity.Book)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.Data(200, "application/epub+zip", b)
	})
	r.GET("/all", func(c *gin.Context) {
		var libraryItems = book.GetAllLibraryItems()

		var libraryItemDtos = make([]dto.LibraryItem, len(libraryItems))
		for i, libraryItem := range libraryItems {
			libraryItemDtos[i] = libraryItem.ToDto()
		}
		c.JSON(200, libraryItemDtos)
	})
	if err := r.Run(":8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	os.MkdirAll("upload/ebooks/", os.ModePerm)
	os.MkdirAll("upload/covers/", os.ModePerm)
	//todo maybe in a different place?
	db := db.GetDbConnection()
	// Migrate the schema
	db.AutoMigrate(&book.Book{})
	db.AutoMigrate(&book.Author{})
	db.AutoMigrate(&book.Collection{})
	setupRoutes()
}
