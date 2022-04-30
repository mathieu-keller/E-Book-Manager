package epubReader

import (
	"archive/zip"
)

func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}
	bk := Book{Fd: fd}
	err = bk.ReadXML("META-INF/container.xml", &bk.Container)
	if err != nil {
		return nil, err
	}
	err = bk.ReadXML(bk.Container.Rootfile.Path, &bk.Opf)
	if err != nil {
		return nil, err
	}
	return &bk, nil
}
