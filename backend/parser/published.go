package parser

import (
	epub2 "e-book-manager/epub"
	"errors"
	"time"
)

func GetDate(bookFile *epub2.Book, e *ParseError) (*time.Time, error) {
	dateField := bookFile.Opf.Metadata.Date
	if len(dateField) > 1 {
		e.Published = "to many dates"
		return nil, errors.New("multi date not supported!")
	} else if len(dateField) == 0 {
		e.Published = "zero dates"
		return nil, errors.New("multi date not supported!")
	}
	var date, err = time.Parse("2006-01-02", dateField[0].Data)
	if err != nil {
		date, err = time.Parse("2006-01-02T15:04:05Z07:00", dateField[0].Data)
		if err != nil {
			e.Published = err.Error()
		}
	}
	return &date, err
}
