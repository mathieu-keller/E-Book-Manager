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

func GetCover(coverId string, bookFile *epub2.Book, bookName string) (string, error) {
	var href = ""
	var imgTyp = ""
	if coverId != "" {
		for _, mani := range bookFile.Opf.Manifest {
			if mani.ID == coverId {
				href = mani.Href
				imgTyp = mani.MediaType
				break
			}
		}
	} else {
		for _, mani := range bookFile.Opf.Manifest {
			if strings.Contains(mani.Href, "cover") || strings.Contains(mani.ID, "cover") {
				if strings.HasSuffix(mani.Href, ".jpg") {
					href = mani.Href
					imgTyp = "image/jpeg"
					break
				}
				if strings.HasSuffix(mani.Href, ".png") {
					href = mani.Href
					imgTyp = "image/png"
					break
				}
				if strings.HasSuffix(mani.Href, ".gif") {
					href = mani.Href
					imgTyp = "image/gif"
					break
				}
			}
		}
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
		err = os.MkdirAll(path, os.ModePerm)
		if imgTyp == "image/jpeg" {
			err = ioutil.WriteFile(path+"cover.jpg", b, fs.ModePerm)
			err = converter.CompressImageResource(path + "cover.jpg")
			if err != nil {
				return "", err
			}
			return path + "cover.jpg", nil
		} else if imgTyp == "image/png" {
			err = ioutil.WriteFile(path+"cover.png", b, fs.ModePerm)
			err = converter.ConvertPngToJpeg(path+"cover.png", path+"cover.jpg")
			if err != nil {
				return "", err
			}
			return path + "cover.jpg", nil
		} else if imgTyp == "image/gif" {
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
