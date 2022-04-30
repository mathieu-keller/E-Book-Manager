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
	mt, err := bk.ReadBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.ReadXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.ReadXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	if bk.Opf.Manifest != nil && bk.Opf.Manifest.Item != nil {
		for _, mf := range *bk.Opf.Manifest.Item {
			if mf.ID == bk.Opf.Spine.Toc {
				err = bk.ReadXML(bk.Filename(mf.Href), &bk.Ncx)
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
