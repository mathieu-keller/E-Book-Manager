package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"os"
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
	newZipFile, err := os.Create(fn)
	if err != nil {
		return nil, err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	b, err := xml.Marshal(bk.Opf)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.WriteFile("test.xml", b, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(bk.Container.Rootfile.Path)
	// Add files to zip
	for _, file := range fd.File {
		if bk.Container.Rootfile.Path == file.Name {
			io, err := zipWriter.Create(bk.Container.Rootfile.Path)
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
	return &bk, nil
}
