package resources

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func InitBookApi(compress *gin.RouterGroup, group *gin.RouterGroup) {
	compressGroup := compress.Group("/book")
	defaultGroup := group.Group("/book")
	compressGroup.GET("/", func(c *gin.Context) {
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

		var books = db.SearchBooks(search, int(page))
		bookDtos := make([]dto.Book, len(books))
		for i, b := range books {
			bookDtos[i] = b.ToDto()
		}
		c.JSON(200, bookDtos)
	})
	compressGroup.GET("/:title", func(c *gin.Context) {
		title := c.Param("title")
		entity := db.GetBookByTitle(title)
		c.JSON(200, entity.ToDto())
	})
	defaultGroup.GET("/download/:id", func(c *gin.Context) {
		id := c.Param("id")
		bookEntity := db.GetBookById(id)
		c.FileAttachment(bookEntity.BookPath, bookEntity.Title+".epub")
	})
}
