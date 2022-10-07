package parser

import (
	"e-book-manager/epub/epubReader"
	"errors"
)

func GetTitle(book *epubReader.Book) (string, error) {
	if book.Opf.Metadata.Title == nil {
		return "", errors.New("no title found")
	}
	titles := *book.Opf.Metadata.Title
	if len(titles) == 0 {
		return "", errors.New("no title found")
	}
	if len(titles) > 1 {
		return "", errors.New("to many title found")
	}
	return titles[0].Text, nil
}
