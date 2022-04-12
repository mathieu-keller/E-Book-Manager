package parser

import "e-book-manager/epub"

func GetSubject(book *epub.Book, e *ParseError) string {
	var subject = book.Opf.Metadata.Subject
	if len(subject) > 1 {
		e.Subject = "to many subjects"
		return subject[0]
	}
	if len(subject) == 0 {
		e.Subject = "no subjects"
		return ""
	}
	return subject[0]
}
