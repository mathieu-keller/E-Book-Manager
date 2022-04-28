package parser

import (
	"e-book-manager/converter"
	epub2 "e-book-manager/epub"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func getHtmlTag(source string, start string, end string) string {
	imgLoc := strings.Index(source, start)
	prefixRem := source[imgLoc+len(start):]
	endImgLoc := strings.Index(prefixRem, end)
	return prefixRem[:endImgLoc]
}

func GetCover(coverId string, bookFile *epub2.Book, bookName string) (*string, error) {
	if bookFile.Opf.Manifest == nil || bookFile.Opf.Manifest.Item == nil {
		return nil, nil
	}
	var href = ""
	var imgTyp = ""
	if coverId != "" {
		for _, mani := range *bookFile.Opf.Manifest.Item {
			if mani.ID == coverId {
				href = mani.Href
				imgTyp = mani.MediaType
				break
			}
		}
	} else {
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
	if imgTyp == "application/xhtml+xml" {
		readedFile, err := bookFile.Open(href)
		if err != nil {
			return nil, err
		}
		defer readedFile.Close()
		b, err := ioutil.ReadAll(readedFile)
		xhtmlString := string(b)
		image := ""
		if strings.Contains(xhtmlString, "<image") {
			imageTag := getHtmlTag(xhtmlString, "<image", "/>")
			image = getHtmlTag(imageTag, "href=\"", "\"")
		} else if strings.Contains(xhtmlString, "<img") {
			imageTag := getHtmlTag(xhtmlString, "<img", "/>")
			image = getHtmlTag(imageTag, "src=\"", "\"")
		}
		href = image
	}
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
		var path = "upload/covers/" + bookName + "/"
		os.MkdirAll(path, os.ModePerm)
		if strings.HasSuffix(href, ".jpg") || strings.HasSuffix(href, ".jpeg") {
			return saveAndConvertCover(path, b, ".jpg")
		} else if strings.HasSuffix(href, ".png") {
			return saveAndConvertCover(path, b, ".png")
		} else if strings.HasSuffix(href, ".gif") {
			return saveAndConvertCover(path, b, ".gif")
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, errors.New("cover not found")
}

func saveAndConvertCover(path string, b []byte, fileEnding string) (*string, error) {
	err := ioutil.WriteFile(path+"cover"+fileEnding, b, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	if fileEnding == ".jpg" {
		err = converter.CompressImageResource(path + "cover" + fileEnding)
	} else if fileEnding == ".png" {
		err = converter.ConvertPngToJpeg(path+"cover"+fileEnding, path+"cover.jpg")
	} else if fileEnding == ".gif" {
		err = converter.ConvertGifToJpeg(path+"cover"+fileEnding, path+"cover.jpg")
	}
	if err != nil {
		return nil, err
	}
	file := path + "cover.jpg"
	return &file, nil
}
