package parser

import (
	"e-book-manager/epub"
	"strconv"
	"strings"
)

func GetCollectionIndex(metaData epub.Metadata) *uint {
	if metaData.Meta == nil {
		return nil
	}
	for _, meta := range *metaData.Meta {
		if strings.HasSuffix(meta.Name, "series_index") {
			index, _ := strconv.ParseFloat(strings.TrimSpace(meta.Content), 8)
			i := uint(index)
			return &i
		}
	}
	return nil
}
