package resources

import (
	"archive/zip"
	"bytes"
	"e-book-manager/parser"
	"github.com/gin-gonic/gin"
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
			file, err := fileHeader.Open()
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			binaryFile := make([]byte, fileHeader.Size)
			fileLength, err := file.Read(binaryFile)
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(fileLength))
			if err != nil {
				fileErrors += "Error: Book " + fileHeader.Filename + ": " + err.Error() + "\n"
				continue
			}
			err = parser.UploadFile(zipReader, fileHeader.Filename)
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
}
