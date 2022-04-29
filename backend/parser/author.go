package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"gorm.io/gorm"
	"strings"
)

func GetAuthor(metaData epub.Metadata, metaIdMap map[string]map[string]epub.Meta, tx *gorm.DB) []*book.Author {
	if metaData.Creator == nil {
		return nil
	}
	var authors = make([]string, 0)
	for _, creator := range *metaData.Creator {
		author := strings.TrimSpace(creator.Text)
		if author != "" {
			if creator.ID == "" || metaIdMap["#"+creator.ID] == nil {
				authors = append(authors, author)
			} else if metaIdMap["#"+creator.ID]["role"].Text == "aut" {
				authors = append(authors, author)
			}
		}
	}
	if len(authors) == 0 {
		return nil
	}
	return createAuth(authors, tx)
}

func createAuth(authorNames []string, tx *gorm.DB) []*book.Author {
	var authors = make([]*book.Author, 0)
	for _, authorName := range authorNames {
		var author = book.GetAuthorByName(authorName, tx)
		if author.Name == "" {
			author.Name = authorName
			author.Create(tx)
		}
		authors = append(authors, &author)
	}

	return authors
}
