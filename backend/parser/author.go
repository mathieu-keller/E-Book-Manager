package parser

import (
	"e-book-manager/db"
	"errors"
	"github.com/mathieu-keller/epub-parser"
	"gorm.io/gorm"
	"strings"
)

func GetAuthor(metadata epub.Metadata, tx *gorm.DB) ([]*db.Author, error) {
	creators, err := getCreator(metadata)
	if err != nil {
		return nil, err
	}
	var authors []*db.Author
	for _, creator := range creators {
		var author = db.GetAuthorByName(creator, tx)
		if author.Name == "" {
			author.Name = creator
			author.Create(tx)
		}
		authors = append(authors, &author)
	}
	return authors, nil
}

func getCreator(metadata epub.Metadata) ([]string, error) {
	if metadata.Creator == nil {
		return nil, errors.New("no creator found")
	}
	creatorsFromMetadata := *metadata.Creator
	if len(creatorsFromMetadata) == 0 {
		return nil, errors.New("no creator found")
	}
	var creators []string
	for _, creator := range creatorsFromMetadata {
		if creator.Role == "aut" || creator.Role == "" {
			creators = append(creators, strings.TrimSpace(creator.Text))
		}
	}
	return creators, nil
}
