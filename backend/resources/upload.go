package resources

import (
	"e-book-manager/parser"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			tmpFileName := uuid.New().String() + fileHeader.Filename
			err := c.SaveUploadedFile(fileHeader, "upload/tmp/"+tmpFileName)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			err = parser.UploadFile(tmpFileName, fileHeader.Filename)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				os.Remove("upload/tmp/" + tmpFileName)
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
