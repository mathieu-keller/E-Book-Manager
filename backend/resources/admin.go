package resources

import (
	"e-book-manager/epub/convert"
	"e-book-manager/parser"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func InitAdminApi(r *gin.RouterGroup) {
	group := r.Group("/admin")
	group.GET("/reimport", func(c *gin.Context) {
		var files []string
		root := "upload/tmp/"
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		for i, file := range files {
			fmt.Println("scan " + strconv.Itoa(i+1) + "/" + strconv.Itoa(len(files)) + " -> " + file)
			bookFile, err := convert.Open(file)
			if err != nil {
				os.Remove(file)
				fmt.Println(err.Error())
				continue
			}
			name := strings.ReplaceAll(file, root, "")
			err = parser.ParseBook(bookFile, root, name)
			bookFile.Close()
			if err != nil {
				os.Remove(file)
				fmt.Println(err.Error())
			}
		}
	})
}
