package parser

import (
	"e-book-manager/db"
	"e-book-manager/epub/epubReader"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func GetAuthor(metadata epubReader.Metadata, tx *gorm.DB) (*db.Author, error) {
	creator, err := getCreator(metadata)
	if err != nil {
		return nil, err
	}
	var author = db.GetAuthorByName(creator, tx)
	if author.Name == "" {
		author.Name = creator
		author.Create(tx)
	}
	return &author, nil
}

func getCreator(metadata epubReader.Metadata) (string, error) {
	if metadata.Creator == nil {
		return "", errors.New("no creator found")
	}
	creators := *metadata.Creator
	if len(creators) == 0 {
		return "", errors.New("no creator found")
	}
	if len(creators) > 1 {
		return "", errors.New("to many creator found")
	}
	return strings.TrimSpace(creators[0].Text), nil
}
