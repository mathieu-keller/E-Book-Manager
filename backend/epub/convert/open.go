package convert

import (
	"archive/zip"
	"e-book-manager/book"
	"e-book-manager/epub"
	"e-book-manager/epub/mash"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

func Open(fn string) (*epub.Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk := epub.Book{Fd: fd}
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

func CopyZip(book epub.Book, entity book.Book) {
	filePath := "upload/ebooks/" + strconv.Itoa(int(entity.ID)) + "-" + entity.Title + ".epub"
	newZipFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	readBook := book.Opf
	b, err := xml.MarshalIndent(ToWriteBook(readBook), "", "    ")
	b = []byte(xml.Header + string(b))
	if err != nil {
		fmt.Println(err.Error())
	}
	// Add files to zip
	for _, file := range book.Fd.File {
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
	entity.Book = filePath
	entity.Persist()
}
func ToWriteBook(p epub.Package) mash.Package {
	opfPackage := mash.Package{
		XMLName:          p.XMLName,
		Version:          p.Version,
		UniqueIdentifier: p.UniqueIdentifier,
		ID:               p.ID,
		Prefix:           p.Prefix,
		Lang:             p.Lang,
		Dir:              p.Dir,
		Metadata:         nil,
		Manifest:         nil,
		Spine:            nil,
		Guide:            nil,
		Bindings:         nil,
	}
	if p.Metadata != nil {
		metadata := mash.Metadata{
			XMLName: p.Metadata.XMLName,
			ID:      p.ID,
			Lang:    p.Lang,
			Opf:     "http://www.idpf.org/2007/opf",
			Dc:      "http://purl.org/dc/elements/1.1/",
			Dcterms: "http://purl.org/dc/terms/",
			Xsi:     "http://www.w3.org/2001/XMLSchema-instance",
			Dir:     p.Dir,
		}
		if p.Metadata.Identifier != nil {
			identifiers := make([]mash.Identifier, len(*p.Metadata.Identifier))
			for i, identifier := range *p.Metadata.Identifier {
				identifiers[i] = mash.Identifier(identifier)
			}
			metadata.Identifier = &identifiers
		}
		if p.Metadata.Title != nil {
			titles := make([]mash.Title, len(*p.Metadata.Title))
			for i, title := range *p.Metadata.Title {
				titles[i] = mash.Title(title)
			}
			metadata.Title = &titles
		}
		if p.Metadata.Language != nil {
			langs := make([]mash.Language, len(*p.Metadata.Language))
			for i, lang := range *p.Metadata.Language {
				langs[i] = mash.Language(lang)
			}
			metadata.Language = &langs
		}
		if p.Metadata.Date != nil {
			dates := make([]mash.Date, len(*p.Metadata.Date))
			for i, date := range *p.Metadata.Date {
				dates[i] = mash.Date(date)
			}
			metadata.Date = &dates
		}
		if p.Metadata.Date != nil {
			dates := make([]mash.Date, len(*p.Metadata.Date))
			for i, date := range *p.Metadata.Date {
				dates[i] = mash.Date(date)
			}
			metadata.Date = &dates
		}
		if p.Metadata.Source != nil {
			sources := make([]mash.Source, len(*p.Metadata.Source))
			for i, source := range *p.Metadata.Source {
				sources[i] = mash.Source(source)
			}
			metadata.Source = &sources
		}
		if p.Metadata.Type != nil {
			types := make([]mash.Type, len(*p.Metadata.Type))
			for i, metaType := range *p.Metadata.Type {
				types[i] = mash.Type(metaType)
			}
			metadata.Type = &types
		}
		if p.Metadata.Format != nil {
			formats := make([]mash.Format, len(*p.Metadata.Format))
			for i, format := range *p.Metadata.Format {
				formats[i] = mash.Format(format)
			}
			metadata.Format = &formats
		}
		if p.Metadata.Creator != nil {
			creators := make([]mash.Creator, len(*p.Metadata.Creator))
			for i, creator := range *p.Metadata.Creator {
				creators[i] = mash.Creator(creator)
			}
			metadata.Creator = &creators
		}
		if p.Metadata.Subject != nil {
			subjects := make([]mash.Subject, len(*p.Metadata.Subject))
			for i, subject := range *p.Metadata.Subject {
				subjects[i] = mash.Subject(subject)
			}
			metadata.Subject = &subjects
		}
		if p.Metadata.Description != nil {
			descriptions := make([]mash.Description, len(*p.Metadata.Description))
			for i, description := range *p.Metadata.Description {
				descriptions[i] = mash.Description(description)
			}
			metadata.Description = &descriptions
		}
		if p.Metadata.Publisher != nil {
			publishers := make([]mash.Publisher, len(*p.Metadata.Publisher))
			for i, publisher := range *p.Metadata.Publisher {
				publishers[i] = mash.Publisher(publisher)
			}
			metadata.Publisher = &publishers
		}
		if p.Metadata.Contributor != nil {
			contributors := make([]mash.Contributor, len(*p.Metadata.Contributor))
			for i, contributor := range *p.Metadata.Contributor {
				contributors[i] = mash.Contributor(contributor)
			}
			metadata.Contributor = &contributors
		}
		if p.Metadata.Relation != nil {
			relations := make([]mash.Relation, len(*p.Metadata.Relation))
			for i, relation := range *p.Metadata.Relation {
				relations[i] = mash.Relation(relation)
			}
			metadata.Relation = &relations
		}
		if p.Metadata.Coverage != nil {
			coverages := make([]mash.Coverage, len(*p.Metadata.Coverage))
			for i, coverage := range *p.Metadata.Coverage {
				coverages[i] = mash.Coverage(coverage)
			}
			metadata.Coverage = &coverages
		}
		if p.Metadata.Rights != nil {
			rights := make([]mash.Rights, len(*p.Metadata.Rights))
			for i, right := range *p.Metadata.Rights {
				rights[i] = mash.Rights(right)
			}
			metadata.Rights = &rights
		}
		if p.Metadata.Meta != nil {
			metas := make([]mash.Meta, len(*p.Metadata.Meta))
			for i, meta := range *p.Metadata.Meta {
				metas[i] = mash.Meta(meta)
			}
			metadata.Meta = &metas
		}
		if p.Metadata.Link != nil {
			links := make([]mash.Link, len(*p.Metadata.Link))
			for i, link := range *p.Metadata.Link {
				links[i] = mash.Link(link)
			}
			metadata.Link = &links
		}
		opfPackage.Metadata = &metadata
	}
	if p.Manifest != nil {
		manifest := mash.Manifest{
			ID:   p.Manifest.ID,
			Item: nil,
		}
		if p.Manifest.Item != nil {
			items := make([]mash.Item, len(*p.Manifest.Item))
			for i, item := range *p.Manifest.Item {
				items[i] = mash.Item(item)
			}
			manifest.Item = &items
		}
		opfPackage.Manifest = &manifest
	}
	if p.Spine != nil {
		spine := mash.Spine{
			ID:                       p.Spine.ID,
			Toc:                      p.Spine.Toc,
			PageProgressionDirection: p.Spine.PageProgressionDirection,
			Itemref:                  nil,
		}
		if p.Spine.Itemref != nil {
			itemRefs := make([]mash.Itemref, len(*p.Spine.Itemref))
			for i, itemRef := range *p.Spine.Itemref {
				itemRefs[i] = mash.Itemref(itemRef)
			}
			spine.Itemref = &itemRefs
		}
		opfPackage.Spine = &spine
	}
	if p.Guide != nil {
		guide := mash.Guide{Reference: nil}
		if p.Guide.Reference != nil {
			references := make([]mash.Reference, len(*p.Guide.Reference))
			for i, reference := range *p.Guide.Reference {
				references[i] = mash.Reference(reference)
			}
			guide.Reference = &references
		}
		opfPackage.Guide = &guide
	}
	if p.Bindings != nil {
		bindings := mash.Bindings{MediaType: nil}
		if p.Bindings.MediaType != nil {
			mediaTypes := make([]mash.MediaType, len(*p.Bindings.MediaType))
			for i, mediaType := range *p.Bindings.MediaType {
				mediaTypes[i] = mash.MediaType(mediaType)
			}
			bindings.MediaType = &mediaTypes
		}
		opfPackage.Bindings = &bindings
	}
	return opfPackage
}
