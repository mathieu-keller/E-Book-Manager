package main

import (
	"e-book-manager/book"
	"e-book-manager/db"
	"e-book-manager/dto"
	"e-book-manager/epub/convert"
	"e-book-manager/parser"
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func uploadFile(fileHeader *multipart.FileHeader) error {
	bookFile, err := convert.Open("upload/tmp/" + fileHeader.Filename)
	if err != nil {
		return err
	}
	defer bookFile.Close()
	err = parser.ParseBook(bookFile, "upload/tmp/", fileHeader.Filename)
	if err != nil {
		os.Remove("upload/tmp/" + fileHeader.Filename)
		return err
	}
	return nil
}

func setupRoutes() {
	r := gin.Default()
	username, userNameSetted := os.LookupEnv("user")
	password, passwordSetted := os.LookupEnv("password")
	compress := r.Group("/")
	compress.Use(gzip.Gzip(gzip.BestCompression))
	var stdApi *gin.RouterGroup = nil
	var defaultAuth *gin.RouterGroup = nil
	if userNameSetted && passwordSetted {
		stdApi = compress.Group("/api", gin.BasicAuth(gin.Accounts{
			username: password,
		},
		))
		defaultAuth = r.Group("/", gin.BasicAuth(gin.Accounts{
			username: password,
		},
		))
	} else {
		stdApi = compress.Group("/api")
		defaultAuth = r.Group("/")
	}
	r.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		static.Serve("/", static.LocalFile("./bundles", true))(c)
	})
	r.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=604800, immutable")
		c.File("./bundles/index.html")
	})
	stdApi.POST("/upload/multi", func(c *gin.Context) {
		files, _ := c.MultipartForm()
		fileErrors := ""
		for _, fileHeader := range files.File["myFiles"] {
			if fileHeader.Header.Get("Content-Type") != "application/epub+zip" {
				fileErrors += "Error: Book " + fileHeader.Filename + ": is not in epub format\n"
				continue
			}
			err := c.SaveUploadedFile(fileHeader, "upload/tmp/"+fileHeader.Filename)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			err = uploadFile(fileHeader)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
		}
		if len(fileErrors) > 0 {
			c.String(400, fileErrors)
		} else {
			c.JSON(200, "Done")
		}
	})
	stdApi.GET("/library/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		entity := book.GetLibraryItemByCollectionId(id)
		c.JSON(200, entity.ToDto())
	})
	stdApi.GET("/book", func(c *gin.Context) {
		queryParam, exist := c.GetQuery("q")
		if !exist {
			c.String(400, "query param q expected")
		}
		pageQuery, exist := c.GetQuery("page")
		if !exist {
			pageQuery = "1"
		}
		page, err := strconv.ParseUint(pageQuery, 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		search := strings.Split(queryParam, " ")

		var books = book.SearchBooks(search, int(page))
		bookDtos := make([]dto.Book, len(books))
		for i, b := range books {
			bookDtos[i] = b.ToDto()
		}
		c.JSON(200, bookDtos)
	})
	stdApi.GET("/book/:title", func(c *gin.Context) {
		title := c.Param("title")
		entity := book.GetBookByTitle(title)
		c.JSON(200, entity.ToDto())
	})
	stdApi.GET("/collection", func(c *gin.Context) {
		title := c.Query("title")
		byName := book.GetCollectionByName(title)
		c.JSON(200, byName.ToDto())
	})
	stdApi.GET("/collection/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		byName := book.GetCollectionById(id)
		c.JSON(200, byName.ToDto())
	})
	defaultAuth.GET("/api/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookEntity := book.GetBookById(id)
		c.FileAttachment(bookEntity.BookPath, bookEntity.Title+".epub")
	})
	stdApi.GET("/all", func(c *gin.Context) {
		pageQuery, exist := c.GetQuery("page")
		if !exist {
			pageQuery = "1"
		}
		page, err := strconv.ParseUint(pageQuery, 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		var libraryItems = book.GetAllLibraryItems(int(page))

		var libraryItemDtos = make([]dto.LibraryItem, len(libraryItems))
		for i, libraryItem := range libraryItems {
			libraryItemDtos[i] = libraryItem.ToDto()
		}
		c.JSON(200, libraryItemDtos)
	})
	stdApi.GET("/reimport", func(c *gin.Context) {
		var files []string
		root := "upload/tmp/"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
		for i, file := range files {
			fmt.Println("scan " + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(files)) + " -> " + file)
			bookFile, err := convert.Open(file)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			name := strings.ReplaceAll(file, root, "")
			err = parser.ParseBook(bookFile, root, name)
			if err != nil {
				bookFile.Close()
				os.Remove(file)
				fmt.Println(err.Error())
			}
			bookFile.Close()
		}
	})
	if err := r.Run(":8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	err := os.MkdirAll("upload/tmp/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = os.MkdirAll("upload/ebooks/", os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	dbCon := db.GetDbConnection()
	err = dbCon.AutoMigrate(&book.Book{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Author{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Subject{})
	if err != nil {
		panic(err.Error())
	}
	err = dbCon.AutoMigrate(&book.Collection{})
	if err != nil {
		panic(err.Error())
	}
	setupRoutes()
}
