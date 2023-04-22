package parser

import (
	"e-book-manager/converter"
	"encoding/base64"
	"errors"
	"github.com/mathieu-keller/epub-parser"
	"io/ioutil"
	"strings"
)

func GetCover(coverId string, bookFile *epub.Book) (*string, error) {
	if bookFile.Opf.Manifest == nil || bookFile.Opf.Manifest.Item == nil {
		return nil, nil
	}
	href := getCoverHref(coverId, bookFile)
	if href != "" {
		readedFile, err := bookFile.Open(href)
		if err != nil {
			return nil, err
		}
		defer readedFile.Close()
		b, err := ioutil.ReadAll(readedFile)
		if err != nil {
			return nil, err
		}
		return saveAndConvertCover(b, href)
	}
	return nil, errors.New("cover not found")
}

func getCoverHref(coverId string, bookFile *epub.Book) string {
	href := ""
	if coverId != "" {
		for _, mani := range *bookFile.Opf.Manifest.Item {
			if mani.ID == coverId {
				href = mani.Href
				break
			}
		}
	} else {
		for _, mani := range *bookFile.Opf.Manifest.Item {
			if mani.Properties == "cover-image" {
				href = mani.Href
				break
			}
		}
		if href == "" {
			for _, mani := range *bookFile.Opf.Manifest.Item {
				if strings.Contains(mani.Href, "cover") || strings.Contains(mani.ID, "cover") {
					if strings.HasSuffix(mani.Href, ".jpg") ||
						strings.HasSuffix(mani.Href, ".png") ||
						strings.HasSuffix(mani.Href, ".gif") {
						href = mani.Href
						break
					}
				}
			}
		}
	}
	return href
}

func saveAndConvertCover(b []byte, href string) (*string, error) {
	resizeImage := []byte{}
	var err error
	if strings.HasSuffix(href, ".svg") {
		image := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(b)
		return &image, nil
	} else if strings.HasSuffix(href, ".jpg") || strings.HasSuffix(href, ".jpeg") {
		resizeImage, err = converter.CompressImageResource(b)
	} else if strings.HasSuffix(href, ".png") {
		resizeImage, err = converter.ConvertPngToJpeg(b)
	} else if strings.HasSuffix(href, ".gif") {
		resizeImage, err = converter.ConvertGifToJpeg(b)
	}
	if err != nil {
		return nil, err
	}
	image := "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(resizeImage)
	return &image, nil
}
