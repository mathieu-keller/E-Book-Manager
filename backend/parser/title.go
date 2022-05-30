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
		if title == "" {
			titles = append(titles, title)
		} else {
			if titleMeta.ID == "" || metaIdMap["#"+titleMeta.ID] == nil ||
				metaIdMap["#"+titleMeta.ID]["title-type"].Text == "" || metaIdMap["#"+titleMeta.ID]["title-type"].Text == "main" {
				if metaIdMap["#"+titleMeta.ID]["file-as"].Text != "" {
					titles = append(titles, metaIdMap["#"+titleMeta.ID]["file-as"].Text)
				} else {
					titles = append(titles, title)
				}

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
