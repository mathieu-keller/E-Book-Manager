package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"strings"
)

func GetSubject(epub *epub.Book) []*book.Subject {
	var subjects = epub.Opf.Metadata.Subject
	subjectEntities := make([]*book.Subject, 0)
	for _, subject := range subjects {
		var trimmedSubject = strings.TrimSpace(subject)
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
