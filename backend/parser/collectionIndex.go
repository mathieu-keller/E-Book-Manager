package parser

import (
	"e-book-manager/epub"
	"strconv"
	"strings"
)

func GetCollectionIndex(bookFile *epub.Book) uint {
	for _, metafield := range bookFile.Opf.Metadata.Meta {
		if strings.HasSuffix(metafield.Name, "series_index") {
			index, _ := strconv.ParseFloat(strings.TrimSpace(metafield.Content), 8)
			return uint(index)
		}
	}
	return 0
}
