package parser

import (
	"errors"
	"github.com/mathieu-keller/epub-parser"
	"strings"
	"time"
)

func GetDate(metaData epub.Metadata) (*time.Time, error) {
	if metaData.Date == nil || len(*metaData.Date) == 0 {
		return nil, errors.New("no date found")
	} else if len(*metaData.Date) > 1 {
		return nil, errors.New("multi date not supported")
	}
	dateField := *metaData.Date
	dateString := strings.TrimSpace(dateField[0].Text)
	var date, err = time.Parse("2006-01-02", dateString)
	if err != nil {
		date, err = time.Parse("2006-01-02T15:04:05Z07:00", dateString)
	}
	return &date, err
}
