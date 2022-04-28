package parser

import (
	"e-book-manager/epub"
	"errors"
	"strings"
	"time"
)

func GetDate(metaData epub.Metadata) (*time.Time, error) {
	dateField := *metaData.Date
	if dateField == nil || len(dateField) == 0 {
		return nil, errors.New("no date found")
	} else if len(dateField) > 1 {
		return nil, errors.New("multi date not supported")
	}
	dateString := strings.TrimSpace(dateField[0].Text)
	var date, err = time.Parse("2006-01-02", dateString)
	if err != nil {
		date, err = time.Parse("2006-01-02T15:04:05Z07:00", dateString)
	}
	return &date, err
}
