package parser

import (
	epub2 "e-book-manager/epub"
	"errors"
	"strings"
	"time"
)

func GetDate(bookFile *epub2.Book) (*time.Time, error) {
	dateField := bookFile.Opf.Metadata.Date
	if len(dateField) > 1 {
		return nil, errors.New("multi date not supported")
	} else if len(dateField) == 0 {
		return nil, errors.New("no date found")
	}
	dateString := strings.TrimSpace(dateField[0].Data)
	var date, err = time.Parse("2006-01-02", dateString)
	if err != nil {
		date, err = time.Parse("2006-01-02T15:04:05Z07:00", dateString)
	}
	return &date, err
}
