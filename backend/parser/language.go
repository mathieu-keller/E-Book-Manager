package parser

import (
	"e-book-manager/epub"
	"errors"
)

func GetLanguage(metaData epub.Metadata) (string, error) {
	lang := *metaData.Language
	if lang == nil || len(lang) == 0 {
		return "", errors.New("lang not found")
	} else if len(lang) > 1 {
		return "", errors.New("multi lang not supported")
	}
	return lang[0].Text, nil
}
