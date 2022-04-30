package epubWriter

import (
	"archive/zip"
	"e-book-manager/epub/epubReader"
	"encoding/xml"
	"os"
)

func CopyZip(book *epubReader.Book, filePath string) error {
	newZipFile, err := os.Create(filePath + "book.epub")
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
func ToWriteBook(p epubReader.Package) Package {
	opfPackage := Package{
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
		manifest := Manifest{
			ID:   p.Manifest.ID,
			Item: nil,
		}
		if p.Manifest.Item != nil {
			items := make([]Item, len(*p.Manifest.Item))
			for i, item := range *p.Manifest.Item {
				items[i] = Item(item)
			}
			manifest.Item = &items
		}
		opfPackage.Manifest = &manifest
	}
	if p.Spine != nil {
		spine := Spine{
			ID:                       p.Spine.ID,
			Toc:                      p.Spine.Toc,
			PageProgressionDirection: p.Spine.PageProgressionDirection,
			Itemref:                  nil,
		}
		if p.Spine.Itemref != nil {
			itemRefs := make([]Itemref, len(*p.Spine.Itemref))
			for i, itemRef := range *p.Spine.Itemref {
				itemRefs[i] = Itemref(itemRef)
			}
			spine.Itemref = &itemRefs
		}
		opfPackage.Spine = &spine
	}
	if p.Guide != nil {
		guide := Guide{Reference: nil}
		if p.Guide.Reference != nil {
			references := make([]Reference, len(*p.Guide.Reference))
			for i, reference := range *p.Guide.Reference {
				references[i] = Reference(reference)
			}
			guide.Reference = &references
		}
		opfPackage.Guide = &guide
	}
	if p.Bindings != nil {
		bindings := Bindings{MediaType: nil}
		if p.Bindings.MediaType != nil {
			mediaTypes := make([]MediaType, len(*p.Bindings.MediaType))
			for i, mediaType := range *p.Bindings.MediaType {
				mediaTypes[i] = MediaType(mediaType)
			}
			bindings.MediaType = &mediaTypes
		}
		opfPackage.Bindings = &bindings
	}
	if p.Collection != nil {
		collections := make([]Collection, len(*p.Collection))
		for i, collection := range *p.Collection {
			collections[i] = getCollection(collection)
		}
	}
	return opfPackage
}

func getCollection(c epubReader.Collection) Collection {
	collection := Collection{
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
		collections := make([]Collection, len(*c.Collections))
		for i, collection := range *c.Collections {
			collections[i] = getCollection(collection)
		}
	}
	return collection
}

func getMetadata(m *epubReader.Metadata) *Metadata {
	metadata := Metadata{
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
		identifiers := make([]Identifier, len(*m.Identifier))
		for i, identifier := range *m.Identifier {
			identifiers[i] = Identifier(identifier)
		}
		metadata.Identifier = &identifiers
	}
	if m.Title != nil {
		titles := make([]Title, len(*m.Title))
		for i, title := range *m.Title {
			titles[i] = Title(title)
		}
		metadata.Title = &titles
	}
	if m.Language != nil {
		langs := make([]Language, len(*m.Language))
		for i, lang := range *m.Language {
			langs[i] = Language(lang)
		}
		metadata.Language = &langs
	}
	if m.Date != nil {
		dates := make([]Date, len(*m.Date))
		for i, date := range *m.Date {
			dates[i] = Date(date)
		}
		metadata.Date = &dates
	}
	if m.Date != nil {
		dates := make([]Date, len(*m.Date))
		for i, date := range *m.Date {
			dates[i] = Date(date)
		}
		metadata.Date = &dates
	}
	if m.Source != nil {
		sources := make([]Source, len(*m.Source))
		for i, source := range *m.Source {
			sources[i] = Source(source)
		}
		metadata.Source = &sources
	}
	if m.Type != nil {
		types := make([]Type, len(*m.Type))
		for i, metaType := range *m.Type {
			types[i] = Type(metaType)
		}
		metadata.Type = &types
	}
	if m.Format != nil {
		formats := make([]Format, len(*m.Format))
		for i, format := range *m.Format {
			formats[i] = Format(format)
		}
		metadata.Format = &formats
	}
	if m.Creator != nil {
		creators := make([]Creator, len(*m.Creator))
		for i, creator := range *m.Creator {
			creators[i] = Creator(creator)
		}
		metadata.Creator = &creators
	}
	if m.Subject != nil {
		subjects := make([]Subject, len(*m.Subject))
		for i, subject := range *m.Subject {
			subjects[i] = Subject(subject)
		}
		metadata.Subject = &subjects
	}
	if m.Description != nil {
		descriptions := make([]Description, len(*m.Description))
		for i, description := range *m.Description {
			descriptions[i] = Description(description)
		}
		metadata.Description = &descriptions
	}
	if m.Publisher != nil {
		publishers := make([]Publisher, len(*m.Publisher))
		for i, publisher := range *m.Publisher {
			publishers[i] = Publisher(publisher)
		}
		metadata.Publisher = &publishers
	}
	if m.Contributor != nil {
		contributors := make([]Contributor, len(*m.Contributor))
		for i, contributor := range *m.Contributor {
			contributors[i] = Contributor(contributor)
		}
		metadata.Contributor = &contributors
	}
	if m.Relation != nil {
		relations := make([]Relation, len(*m.Relation))
		for i, relation := range *m.Relation {
			relations[i] = Relation(relation)
		}
		metadata.Relation = &relations
	}
	if m.Coverage != nil {
		coverages := make([]Coverage, len(*m.Coverage))
		for i, coverage := range *m.Coverage {
			coverages[i] = Coverage(coverage)
		}
		metadata.Coverage = &coverages
	}
	if m.Rights != nil {
		rights := make([]Rights, len(*m.Rights))
		for i, right := range *m.Rights {
			rights[i] = Rights(right)
		}
		metadata.Rights = &rights
	}
	if m.Meta != nil {
		metas := make([]Meta, len(*m.Meta))
		for i, meta := range *m.Meta {
			metas[i] = Meta(meta)
		}
		metadata.Meta = &metas
	}
	if m.Link != nil {
		metadata.Link = getLinks(m.Link)
	}
	return &metadata
}

func getLinks(link *[]epubReader.Link) *[]Link {
	links := make([]Link, len(*link))
	for i, link := range *link {
		links[i] = Link(link)
	}
	return &links
}
