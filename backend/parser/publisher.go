package parser

import (
	"e-book-manager/epub"
	"errors"
)

func GetPublisher(metaData epub.Metadata) (*string, error) {
	pub := *metaData.Publisher
	if pub == nil || len(pub) == 0 {
		return nil, errors.New("no publisher found")
	} else if len(pub) > 1 {
		return nil, errors.New("multi publisher not supported")
	}
	return &pub[0].Text, nil
}
