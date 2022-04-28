package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"strings"
)

func GetSubject(metaData epub.Metadata) []*book.Subject {
	if metaData.Subject == nil {
		return nil
	}
	var subjects = *metaData.Subject
	subjectEntities := make([]*book.Subject, 0)
	for _, subject := range subjects {
		var trimmedSubject = strings.TrimSpace(subject.Text)
		if trimmedSubject != "" {
			var entity = book.GetSubjectByName(trimmedSubject)
			if entity.Name == "" {
				entity.Name = trimmedSubject
				entity.Persist()
			}
			subjectEntities = append(subjectEntities, &entity)
		}
	}

	return subjectEntities
}
