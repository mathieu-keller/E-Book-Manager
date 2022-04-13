package parser

import (
	"e-book-manager/epub"
	"strings"
)

func GetTitle(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Metafield, e *ParseError) string {
	var titles = make([]string, 0)
	for _, titleMeta := range bookFile.Opf.Metadata.Title {
		var title = strings.TrimSpace(titleMeta.Data)
		if title != "" {
			if titleMeta.ID == "" || metaIdMap["#"+titleMeta.ID] == nil {
				titles = append(titles, title)
			} else if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "main" {
				titles = append(titles, title)
			}
		}
	}
	if len(titles) > 1 {
		e.Title = "to many titles"
		return titles[0]
	}
	if len(titles) == 0 {
		e.Title = "not found!"
		return ""
	}
	return titles[0]
}
