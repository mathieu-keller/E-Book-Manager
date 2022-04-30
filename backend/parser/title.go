package parser

import (
	"e-book-manager/epub/epubReader"
	"strings"
)

func GetTitle(metaData epubReader.Metadata, metaIdMap map[string]map[string]epubReader.Meta) string {
	if metaData.Title == nil {
		return ""
	}
	var titles = make([]string, 0)
	for _, titleMeta := range *metaData.Title {
		var title = strings.TrimSpace(titleMeta.Text)
		if title != "" {
			if titleMeta.ID == "" || metaIdMap["#"+titleMeta.ID] == nil {
				titles = append(titles, title)
			} else if metaIdMap["#"+titleMeta.ID]["title-type"].Text == "main" {
				titles = append(titles, title)
			}
		}
	}
	if len(titles) > 1 {
		return titles[0]
	}
	if len(titles) == 0 {
		return ""
	}
	return titles[0]
}
