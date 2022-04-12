package main

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/dto"
	epub2 "e-book-manager/epub"
	"e-book-manager/parser"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"strconv"
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
	parseError := parser.ParseError{}
	bookEntity := book.Book{}
	bookEntity.Author = parser.GetAuthor(bookFile, metaIdMap, &parseError)
	var date, err = parser.GetDate(bookFile, &parseError)
	if err == nil {
		bookEntity.Published = *date
	}
	bookEntity.Publisher, _ = parser.GetPublisher(bookFile, &parseError)
	bookEntity.Language, _ = parser.GetLanguage(bookFile, &parseError)
	bookEntity.Name = parser.GetTitle(bookFile, metaIdMap, &parseError)
	bookEntity.CollectionId = parser.GetCollection(bookFile, metaIdMap, &parseError)
	bookEntity.Cover, _ = parser.GetCover(coverId, bookFile, bookEntity.Name, &parseError)
	bookEntity.Subject = parser.GetSubject(bookFile, &parseError)
	bookEntity.Book = path
	bookEntity.Persist()
	parseError.Book = bookEntity.ID
	parseError.Persist()
	/*s := strings.Split(bookFile.Container.Rootfile.Path, "/")
	r, err := bookFile.Open(s[len(s)-1])
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	defer r.Close()
	os.MkdirAll("raw/"+strconv.Itoa(int(bookEntity.ID))+"/", os.ModePerm)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	os.WriteFile("raw/"+strconv.Itoa(int(bookEntity.ID))+"/raw.xml", b, os.ModePerm)
	*/
	return &bookEntity, nil
}

func setupRoutes() {
	r := gin.Default()
	r.POST("/upload/multi", func(c *gin.Context) {
		files, _ := c.MultipartForm()

		for i, fileHeader := range files.File["myFiles"] {
			fmt.Println(i)
			fmt.Println(fileHeader.Filename)
			os.MkdirAll("upload/ebooks/", os.ModePerm)
			c.SaveUploadedFile(fileHeader, "upload/ebooks/"+fileHeader.Filename)
			uploadFile(fileHeader)
		}
		c.String(200, "Done")
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("myFile")
		err := os.MkdirAll("upload/ebooks/", os.ModePerm)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		err = c.SaveUploadedFile(file, "upload/ebooks/"+file.Filename)
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
	r.GET("/", func(c *gin.Context) {
		file, err := os.ReadFile("index.html")
		if err != nil {
			c.String(500, err.Error())
		}
		c.Data(200, "text/html; charset=utf-8", file)
	})
	r.GET("/library/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		entity := book.GetLibraryItemByCollectionId(id)
		c.JSON(200, entity.ToDto())
	})
	r.GET("/book/:title", func(c *gin.Context) {
		title := c.Param("title")
		entity := book.GetBookByTitle(title)
		c.JSON(200, entity.ToDto())
	})
	r.GET("/collection", func(c *gin.Context) {
		name := c.Query("name")
		byName := book.GetCollectionByName(name)
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
	r.Run()
}

func main() {
	//todo maybe in a different place?
	db := db.GetDbConnection()
	// Migrate the schema
	db.AutoMigrate(&book.Book{})
	db.AutoMigrate(&book.Author{})
	db.AutoMigrate(&book.Collection{})
	db.AutoMigrate(&parser.ParseError{})
	setupRoutes()
}
