package parser

import (
	"errors"
	"github.com/mathieu-keller/epub-parser"
)

func GetPublisher(metaData epub.Metadata) (*string, error) {
	if metaData.Publisher == nil || len(*metaData.Publisher) == 0 {
		return nil, errors.New("no publisher found")
	} else if len(*metaData.Publisher) > 1 {
		return nil, errors.New("multi publisher not supported")
	}
	pub := *metaData.Publisher
	return &pub[0].Text, nil
}
