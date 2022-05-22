package resources

import (
	"e-book-manager/db"
	"e-book-manager/dto"
	"github.com/gin-gonic/gin"
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
}
