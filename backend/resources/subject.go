package resources

import (
	"archive/zip"
	"bytes"
	"e-book-manager/db"
	"e-book-manager/dto"
	"e-book-manager/epub/epubReader"
	"e-book-manager/epub/epubWriter"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func InitSubjectApi(compress *gin.RouterGroup) {

	compress.GET("subjects", func(c *gin.Context) {
		subjects := db.GetAllSubjects()
		subjectDtos := make([]dto.Subject, len(subjects))
		for i, subject := range subjects {
			subjectDtos[i] = subject.ToDto()
		}
		c.JSON(200, subjectDtos)
	})
	compress.POST("subjects", func(c *gin.Context) {
		book := dto.Book{}
		err := c.BindJSON(&book)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		tx := db.GetDbConnection().Begin()
		subjectEntities := make([]*db.Subject, 0)
		for _, subject := range book.Subjects {
			var trimmedSubject = strings.TrimSpace(subject.Name)
			if trimmedSubject != "" {
				var entity = db.GetSubjectByName(trimmedSubject, tx)
				if entity.Name == "" {
					entity.Name = trimmedSubject
					entity.Persist(tx)
				}
				subjectEntities = append(subjectEntities, &entity)
			}
		}
		file, err := os.Open(book.Book)
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		defer file.Close()
		state, err := file.Stat()
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		binaryFile := make([]byte, state.Size())
		fileLength, err := file.Read(binaryFile)
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(fileLength))
		epub, err := epubReader.Open(zipReader)
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}

		readBook := epub
		subjectNames := make([]epubReader.Subject, len(book.Subjects))
		for i, subject := range book.Subjects {
			subjectNames[i] = epubReader.Subject{Text: subject.Name}
		}
		readBook.Opf.Metadata.Subject = &subjectNames
		epubWriter.CreateZip(readBook, book.Book+"copy")
		err = os.Rename(book.Book, book.Book+"backup")
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		err = os.Rename(book.Book+"copy", book.Book)
		if err != nil {
			os.Rename(book.Book+"backup", book.Book)
			os.Remove(book.Book + "copy")
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		err = os.Remove(book.Book + "backup")
		if err != nil {
			tx.Rollback()
			c.String(500, err.Error())
			return
		}
		tx.Commit()
		subjectDtos := make([]dto.Subject, len(subjectEntities))
		for i, subject := range subjectEntities {
			subjectDtos[i] = subject.ToDto()
		}
		c.JSON(200, subjectDtos)
	})

}
