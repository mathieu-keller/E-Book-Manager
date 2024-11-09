package epubReader

import (
	"archive/zip"
)

func Open(reader *zip.Reader) (*Book, error) {
	bk := Book{Fd: reader}
	err := bk.ReadXML("META-INF/container.xml", &bk.Container)
	if err != nil {
		return nil, err
	}
	err = bk.ReadXML(bk.Container.Rootfile.Path, &bk.Opf)
	if err != nil {
		return nil, err
	}
	return &bk, nil
}
