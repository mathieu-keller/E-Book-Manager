package convert

import (
	"archive/zip"
	"e-book-manager/epub"
	"e-book-manager/epub/mash"
	"encoding/xml"
	"os"
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

func CopyZip(book *epub.Book, filePath string) error {
	newZipFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	readBook := book.Opf
	b, err := xml.MarshalIndent(ToWriteBook(readBook), "", "    ")
	if err != nil {
		return err
	}
	b = []byte(xml.Header + string(b))
	// Add files to zip
	for _, file := range book.Fd.File {
		if file.FileInfo().IsDir() {
			continue
		}
		if book.Container.Rootfile.Path == file.Name {
			io, err := zipWriter.Create(book.Container.Rootfile.Path)
			if err != nil {
				return err
			}
			_, err = io.Write(b)
			if err != nil {
				return err
			}
		} else {
			err = zipWriter.Copy(file)
			if err != nil {
				return err
			}
		}
		err := zipWriter.Flush()
		if err != nil {
			return err
		}
	}
	return newZipFile.Sync()
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
		opfPackage.Metadata = getMetadata(p.Metadata)
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
	if p.Collection != nil {
		collections := make([]mash.Collection, len(*p.Collection))
		for i, collection := range *p.Collection {
			collections[i] = getCollection(collection)
		}
	}
	return opfPackage
}

func getCollection(c epub.Collection) mash.Collection {
	collection := mash.Collection{
		Dir:         c.Dir,
		Id:          c.Id,
		Role:        c.Role,
		Lang:        c.Lang,
		Metadata:    nil,
		Link:        nil,
		Collections: nil,
	}
	if c.Metadata != nil {
		collection.Metadata = getMetadata(c.Metadata)
	}
	if c.Link != nil {
		collection.Link = getLinks(c.Link)
	}
	if c.Collections != nil {
		collections := make([]mash.Collection, len(*c.Collections))
		for i, collection := range *c.Collections {
			collections[i] = getCollection(collection)
		}
	}
	return collection
}

func getMetadata(m *epub.Metadata) *mash.Metadata {
	metadata := mash.Metadata{
		XMLName: m.XMLName,
		ID:      m.ID,
		Lang:    m.Lang,
		Opf:     "http://www.idpf.org/2007/opf",
		Dc:      "http://purl.org/dc/elements/1.1/",
		Dcterms: "http://purl.org/dc/terms/",
		Xsi:     "http://www.w3.org/2001/XMLSchema-instance",
		Dir:     m.Dir,
	}
	if m.Identifier != nil {
		identifiers := make([]mash.Identifier, len(*m.Identifier))
		for i, identifier := range *m.Identifier {
			identifiers[i] = mash.Identifier(identifier)
		}
		metadata.Identifier = &identifiers
	}
	if m.Title != nil {
		titles := make([]mash.Title, len(*m.Title))
		for i, title := range *m.Title {
			titles[i] = mash.Title(title)
		}
		metadata.Title = &titles
	}
	if m.Language != nil {
		langs := make([]mash.Language, len(*m.Language))
		for i, lang := range *m.Language {
			langs[i] = mash.Language(lang)
		}
		metadata.Language = &langs
	}
	if m.Date != nil {
		dates := make([]mash.Date, len(*m.Date))
		for i, date := range *m.Date {
			dates[i] = mash.Date(date)
		}
		metadata.Date = &dates
	}
	if m.Date != nil {
		dates := make([]mash.Date, len(*m.Date))
		for i, date := range *m.Date {
			dates[i] = mash.Date(date)
		}
		metadata.Date = &dates
	}
	if m.Source != nil {
		sources := make([]mash.Source, len(*m.Source))
		for i, source := range *m.Source {
			sources[i] = mash.Source(source)
		}
		metadata.Source = &sources
	}
	if m.Type != nil {
		types := make([]mash.Type, len(*m.Type))
		for i, metaType := range *m.Type {
			types[i] = mash.Type(metaType)
		}
		metadata.Type = &types
	}
	if m.Format != nil {
		formats := make([]mash.Format, len(*m.Format))
		for i, format := range *m.Format {
			formats[i] = mash.Format(format)
		}
		metadata.Format = &formats
	}
	if m.Creator != nil {
		creators := make([]mash.Creator, len(*m.Creator))
		for i, creator := range *m.Creator {
			creators[i] = mash.Creator(creator)
		}
		metadata.Creator = &creators
	}
	if m.Subject != nil {
		subjects := make([]mash.Subject, len(*m.Subject))
		for i, subject := range *m.Subject {
			subjects[i] = mash.Subject(subject)
		}
		metadata.Subject = &subjects
	}
	if m.Description != nil {
		descriptions := make([]mash.Description, len(*m.Description))
		for i, description := range *m.Description {
			descriptions[i] = mash.Description(description)
		}
		metadata.Description = &descriptions
	}
	if m.Publisher != nil {
		publishers := make([]mash.Publisher, len(*m.Publisher))
		for i, publisher := range *m.Publisher {
			publishers[i] = mash.Publisher(publisher)
		}
		metadata.Publisher = &publishers
	}
	if m.Contributor != nil {
		contributors := make([]mash.Contributor, len(*m.Contributor))
		for i, contributor := range *m.Contributor {
			contributors[i] = mash.Contributor(contributor)
		}
		metadata.Contributor = &contributors
	}
	if m.Relation != nil {
		relations := make([]mash.Relation, len(*m.Relation))
		for i, relation := range *m.Relation {
			relations[i] = mash.Relation(relation)
		}
		metadata.Relation = &relations
	}
	if m.Coverage != nil {
		coverages := make([]mash.Coverage, len(*m.Coverage))
		for i, coverage := range *m.Coverage {
			coverages[i] = mash.Coverage(coverage)
		}
		metadata.Coverage = &coverages
	}
	if m.Rights != nil {
		rights := make([]mash.Rights, len(*m.Rights))
		for i, right := range *m.Rights {
			rights[i] = mash.Rights(right)
		}
		metadata.Rights = &rights
	}
	if m.Meta != nil {
		metas := make([]mash.Meta, len(*m.Meta))
		for i, meta := range *m.Meta {
			metas[i] = mash.Meta(meta)
		}
		metadata.Meta = &metas
	}
	if m.Link != nil {
		metadata.Link = getLinks(m.Link)
	}
	return &metadata
}

func getLinks(link *[]epub.Link) *[]mash.Link {
	links := make([]mash.Link, len(*link))
	for i, link := range *link {
		links[i] = mash.Link(link)
	}
	return &links
}
