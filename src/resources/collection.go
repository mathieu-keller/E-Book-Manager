package resources

import (
	"e-book-manager/db"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitCollectionApi(r *gin.RouterGroup) {
	group := r.Group("/collection")
	group.GET("/", func(c *gin.Context) {
		title := c.Query("title")
		byName := db.GetCollectionByName(title)
		c.JSON(200, byName.ToDto())
	})
	group.GET("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 8)
		if err != nil {
			c.String(500, err.Error())
		}
		byName := db.GetCollectionById(id)
		c.JSON(200, byName.ToDto())
	})
}
