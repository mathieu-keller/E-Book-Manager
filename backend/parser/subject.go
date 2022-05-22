package parser

import (
	"e-book-manager/db"
	"e-book-manager/epub/epubReader"
	"gorm.io/gorm"
	"strings"
)

func GetSubject(metaData epubReader.Metadata, tx *gorm.DB) []*db.Subject {
	if metaData.Subject == nil {
		return nil
	}
	var subjects = *metaData.Subject
	subjectEntities := make([]*db.Subject, len(subjects))
	for i, subject := range subjects {
		var trimmedSubject = strings.TrimSpace(subject.Text)
		if trimmedSubject != "" {
			var entity = db.GetSubjectByName(trimmedSubject, tx)
			if entity.Name == "" {
				entity.Name = trimmedSubject
				entity.Persist(tx)
			}
			subjectEntities[i] = &entity
		}
	}

	return subjectEntities
}
