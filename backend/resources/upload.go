package resources

import (
	"e-book-manager/parser"
	"github.com/gin-gonic/gin"
	"os"
)

func InitUploadApi(r *gin.RouterGroup) {
	group := r.Group("/upload")
	group.POST("/multi", func(c *gin.Context) {
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
			err = parser.UploadFile(fileHeader)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				os.Remove("upload/tmp/" + fileHeader.Filename)
				continue
			}
		}
		if len(fileErrors) > 0 {
			c.String(400, fileErrors)
		} else {
			c.JSON(200, "Done")
		}
	})
}