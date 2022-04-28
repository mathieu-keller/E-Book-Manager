package parser

import (
	epub2 "e-book-manager/epub"
	"errors"
)

func GetPublisher(bookFile *epub2.Book) (string, error) {
	pub := *bookFile.Opf.Metadata.Publisher
	if len(pub) > 1 {
		return "nil", errors.New("multi publisher not supported")
	} else if len(pub) == 0 {
		return "", errors.New("multi publisher not supported")
	}
	return pub[0].Text, nil
}
