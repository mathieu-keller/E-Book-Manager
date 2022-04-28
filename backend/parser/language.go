package parser

import (
	epub2 "e-book-manager/epub"
	"errors"
)

func GetLanguage(bookFile *epub2.Book) (string, error) {
	lang := *bookFile.Opf.Metadata.Language
	if len(lang) > 1 {
		return "", errors.New("multi lang not supported")
	} else if len(lang) == 0 {
		return "", errors.New("multi lang not supported")
	}
	return lang[0].Text, nil
}
