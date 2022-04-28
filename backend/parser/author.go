package parser

import (
	"e-book-manager/book"
	epub2 "e-book-manager/epub"
	"strings"
)

func GetAuthor(bookFile *epub2.Book, metaIdMap map[string]map[string]epub2.Meta) []*book.Author {
	var authors = make([]string, 0)
	for _, creator := range *bookFile.Opf.Metadata.Creator {
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
	return createAuth(authors)
}

func createAuth(authorNames []string) []*book.Author {
	var authors = make([]*book.Author, 0)
	for _, authorName := range authorNames {
		var author = book.GetAuthorByName(authorName)
		if author.Name == "" {
			author.Name = authorName
		}
		authors = append(authors, &author)
	}

	return authors
}
