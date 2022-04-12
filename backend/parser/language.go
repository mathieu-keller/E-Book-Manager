package parser

import (
	epub2 "e-book-manager/epub"
	"errors"
)

func GetLanguage(bookFile *epub2.Book, e *ParseError) (string, error) {
	lang := bookFile.Opf.Metadata.Language
	if len(lang) > 1 {
		e.Language = "to many lang"
		return "", errors.New("multi lang not supported!")
	} else if len(lang) == 0 {
		e.Language = "zero langs"
		return "", errors.New("multi lang not supported!")
	}
	return lang[0], nil
}
