package epub

import (
	"archive/zip"
	"e-book-manager/book"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk := Book{fd: fd}
	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	if bk.Opf.Manifest != nil && bk.Opf.Manifest.Item != nil {
		for _, mf := range *bk.Opf.Manifest.Item {
			if mf.ID == bk.Opf.Spine.Toc {
				err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
				break
			}
		}
	}

	if err != nil {
		fd.Close()
		return nil, err
	}
	return &bk, nil
}

func CopyZip(book Book, entity book.Book) {
	newZipFile, err := os.Create("upload/ebooks/" + strconv.Itoa(int(entity.ID)) + "-" + entity.Title + ".epub")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	b, err := xml.Marshal(book.Opf)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Add files to zip
	for _, file := range book.fd.File {
		if file.FileInfo().IsDir() {
			continue
		}
		if book.Container.Rootfile.Path == file.Name {
			io, err := zipWriter.Create(book.Container.Rootfile.Path)
			if err != nil {
				fmt.Println("error create")
				fmt.Println(err.Error())
			}
			_, err = io.Write(b)
			if err != nil {
				fmt.Println("error write")
				fmt.Println(err.Error())
			}
		} else {
			err = zipWriter.Copy(file)
			if err != nil {
				fmt.Println("error copy")
				fmt.Println(err.Error())
			}
		}
	}
}
