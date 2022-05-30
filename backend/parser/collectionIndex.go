package parser

import (
	"e-book-manager/epub/epubReader"
	"strconv"
	"strings"
)

func GetCollectionIndex(metaData epubReader.Metadata) *uint {
	if metaData.Meta == nil {
		return nil
	}
	for _, meta := range *metaData.Meta {
		if strings.HasSuffix(meta.Name, "series_index") || strings.HasSuffix(meta.Property, "group-position") {
			index, _ := strconv.ParseFloat(strings.TrimSpace(meta.Content), 8)
			i := uint(index)
			return &i
		}
	}
	return nil
}
