package resources

import (
	"archive/zip"
	"bytes"
	"e-book-manager/parser"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func InitAdminApi(compress *gin.RouterGroup) {
	defaultGroup := compress.Group("/admin")
	defaultGroup.GET("/server/import", func(c *gin.Context) {
		files, err := ioutil.ReadDir("upload/upload")
		if err != nil {
			c.String(500, err.Error())
			return
		}
		fileErrors := ""
		for _, file := range files {
			if !file.IsDir() {
				f, err := os.Open("upload/upload/" + file.Name())
				if err != nil {
					fileErrors += "Error: Book " + f.Name() + ": " + err.Error() + "\n"
					f.Close()
					continue
				}
				fileBytes := make([]byte, file.Size())
				size, err := f.Read(fileBytes)
				if err != nil {
					fileErrors += "Error: Book " + f.Name() + ": " + err.Error() + "\n"
					f.Close()
					continue
				}
				zipReader, err := zip.NewReader(bytes.NewReader(fileBytes), int64(size))
				if err != nil {
					fileErrors += "Error: Book " + f.Name() + ": " + err.Error() + "\n"
					f.Close()
					continue
				}
				err = parser.UploadFile(zipReader, f.Name())
				if err != nil {
					fileErrors += "Error: Book " + f.Name() + ": " + err.Error() + "\n"
					f.Close()
					continue
				}
				f.Close()
			}
		}
		c.String(200, fileErrors)
	})
}
