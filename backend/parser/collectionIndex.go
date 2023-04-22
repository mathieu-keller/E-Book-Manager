package parser

import (
	"github.com/mathieu-keller/epub-parser"
	"strconv"
	"strings"
)

func GetCollectionIndex(metaData epub.Metadata) *uint {
	if metaData.Meta == nil {
		return nil
	}
	for _, meta := range *metaData.Meta {
		if strings.HasSuffix(meta.Name, "series_index") || strings.HasSuffix(meta.Property, "group-position") {
			index, _ := strconv.ParseFloat(strings.TrimSpace(meta.Text), 8)
			i := uint(index)
			return &i
		}
	}
	return nil
}
