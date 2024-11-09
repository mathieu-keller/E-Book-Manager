package resources

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitLibraryApi(r *gin.RouterGroup) {
	group := r.Group("/library")
	group.GET("/all", func(c *gin.Context) {
		pageQuery, exist := c.GetQuery("page")
		if !exist {
			pageQuery = "1"
		}
		page, err := strconv.ParseUint(pageQuery, 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		var libraryItems = db.GetAllLibraryItems(int(page))

		var libraryItemDtos = make([]dto.LibraryItem, len(libraryItems))
		for i, libraryItem := range libraryItems {
			libraryItemDtos[i] = libraryItem.ToDto()
		}
		c.JSON(200, libraryItemDtos)
	})
}
