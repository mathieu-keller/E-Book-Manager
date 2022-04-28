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

func GetCover(coverId string, bookFile *epub2.Book, bookName string) (string, error) {
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
			return "", err
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
			return "", err
		}
		defer readedFile.Close()
		b, err := ioutil.ReadAll(readedFile)
		if err != nil {
			return "", err
		}
		var path = "upload/covers/" + bookName + "/"
		os.MkdirAll(path, os.ModePerm)
		if strings.HasSuffix(href, ".jpg") || strings.HasSuffix(href, ".jpeg") {
			err = ioutil.WriteFile(path+"cover.jpg", b, fs.ModePerm)
			err = converter.CompressImageResource(path + "cover.jpg")
			if err != nil {
				return "", err
			}
			return path + "cover.jpg", nil
		} else if strings.HasSuffix(href, ".png") {
			err = ioutil.WriteFile(path+"cover.png", b, fs.ModePerm)
			err = converter.ConvertPngToJpeg(path+"cover.png", path+"cover.jpg")
			if err != nil {
				return "", err
			}
			return path + "cover.jpg", nil
		} else if strings.HasSuffix(href, ".gif") {
			err = ioutil.WriteFile(path+"cover.gif", b, fs.ModePerm)
			err = converter.ConvertGifToJpeg(path+"cover.gif", path+"cover.jpg")
			if err != nil {
				return "", err
			}
			return path + "cover.jpg", nil
		}
		if err != nil {
			return "", err
		}
	}
	return "", errors.New("cover not found")
}
