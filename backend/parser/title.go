package parser

import (
	"e-book-manager/epub"
)

func GetTitle(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Metafield, e *ParseError) string {
	var title = make([]string, 0)
	for _, titleMeta := range bookFile.Opf.Metadata.Title {
		if titleMeta.ID == "" || metaIdMap["#"+titleMeta.ID] == nil {
			title = append(title, titleMeta.Data)
		} else if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "main" {
			title = append(title, titleMeta.Data)
		}
	}
	if len(title) > 1 {
		e.Title = "to many titles"
		return title[0]
	}
	if len(title) == 0 {
		e.Title = "not found!"
		return ""
	}
	return title[0]
}
