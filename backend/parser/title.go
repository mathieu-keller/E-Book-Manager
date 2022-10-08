package parser

import (
	"e-book-manager/epub/epubReader"
	"errors"
)

func GetTitle(metadata epubReader.Metadata) (string, error) {
	if metadata.Title == nil {
		return "", errors.New("no title found")
	}
	titles := *metadata.Title
	if len(titles) == 0 {
		return "", errors.New("no title found")
	}
	if len(titles) > 1 {
		return "", errors.New("to many title found")
	}
	return titles[0].Text, nil
}
