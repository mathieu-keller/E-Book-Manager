package mash

import (
	"encoding/xml"
)

type Package struct {
	XMLName          xml.Name  `xml:"http://www.idpf.org/2007/opf package"`
	Version          string    `xml:"version,attr,omitempty"`
	UniqueIdentifier string    `xml:"unique-identifier,attr,omitempty"`
	ID               string    `xml:"id,attr,omitempty"`
	Prefix           string    `xml:"prefix,attr,omitempty"`
	Lang             string    `xml:"lang,attr,omitempty"`
	Dir              string    `xml:"dir,attr,omitempty"`
	Metadata         *Metadata `xml:"metadata,omitempty"`
	Manifest         *Manifest `xml:"manifest,omitempty"`
	Spine            *Spine    `xml:"spine,omitempty"`
	Guide            *Guide    `xml:"guide,omitempty"`
	Bindings         *Bindings `xml:"bindings,omitempty"`
}

type Bindings struct {
	MediaType *[]MediaType `xml:"mediaType,omitempty"`
}

type MediaType struct {
	Text      string `xml:",chardata"`
	MediaType string `xml:"media-type,attr,omitempty"`
	Handler   string `xml:"handler,attr,omitempty"`
}

type Guide struct {
	Reference *[]Reference `xml:"reference,omitempty"`
}

type Reference struct {
	Text  string `xml:",chardata"`
	Href  string `xml:"href,attr,omitempty"`
	Type  string `xml:"type,attr,omitempty"`
	Title string `xml:"title,attr,omitempty"`
}

type Spine struct {
	ID                       string     `xml:"id,attr,omitempty"`
	Toc                      string     `xml:"toc,attr,omitempty"`
	PageProgressionDirection string     `xml:"page-progression-direction,attr,omitempty"`
	Itemref                  *[]Itemref `xml:"itemref,omitempty"`
}

type Itemref struct {
	Text       string `xml:",chardata"`
	Idref      string `xml:"idref,attr,omitempty"`
	Linear     string `xml:"linear,attr,omitempty"`
	ID         string `xml:"id,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
}

type Manifest struct {
	ID   string  `xml:"id,attr,omitempty"`
	Item *[]Item `xml:"item,omitempty"`
}

type Item struct {
	Text         string `xml:",chardata"`
	ID           string `xml:"id,attr,omitempty"`
	Href         string `xml:"href,attr,omitempty"`
	MediaType    string `xml:"media-type,attr,omitempty"`
	Fallback     string `xml:"fallback,attr,omitempty"`
	MediaOverlay string `xml:"media-overlay,attr,omitempty"`
	Properties   string `xml:"properties,attr,omitempty"`
}

type Metadata struct {
	XMLName     xml.Name       `xml:"metadata"`
	ID          string         `xml:"id,attr,omitempty"`
	Lang        string         `xml:"lang,attr,omitempty"`
	Opf         string         `xml:"xmlns:opf,attr"`
	Dc          string         `xml:"xmlns:dc,attr"`
	Dcterms     string         `xml:"xmlns:dcterms,attr"`
	Xsi         string         `xml:"xmlns:xsi,attr"`
	Dir         string         `xml:"dir,attr,omitempty"`
	Identifier  *[]Identifier  `xml:"identifier,omitempty"`
	Title       *[]Title       `xml:"dc:title,omitempty"`
	Language    *[]Language    `xml:"dc:language,omitempty"`
	Date        *[]Date        `xml:"dc:date,omitempty"`
	Source      *[]Source      `xml:"dc:source,omitempty"`
	Type        *[]Type        `xml:"dc:type,omitempty"`
	Format      *[]Format      `xml:"dc:format,omitempty"`
	Creator     *[]Creator     `xml:"dc:creator,omitempty"`
	Subject     *[]Subject     `xml:"dc:subject,omitempty"`
	Description *[]Description `xml:"dc:description,omitempty"`
	Publisher   *[]Publisher   `xml:"dc:publisher,omitempty"`
	Contributor *[]Contributor `xml:"dc:contributor,omitempty"`
	Relation    *[]Relation    `xml:"dc:relation,omitempty"`
	Coverage    *[]Coverage    `xml:"dc:coverage,omitempty"`
	Rights      *[]Rights      `xml:"dc:rights,omitempty"`
	Meta        *[]Meta        `xml:"meta,omitempty"`
	Link        *[]Link        `xml:"dc:link,omitempty"`
}

type Identifier struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Scheme string `xml:"opf:scheme,attr,omitempty"`
}

type Title struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Language struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Date struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Source struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Type struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Format struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
}

type Creator struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Role   string `xml:"opf:role,attr,omitempty"`
	FileAs string `xml:"opf:file-as,attr,omitempty"`
	Lang   string `xml:"lang,attr,omitempty"`
	Dir    string `xml:"dir,attr,omitempty"`
}

type Subject struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Description struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Publisher struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Contributor struct {
	Text   string `xml:",chardata"`
	ID     string `xml:"id,attr,omitempty"`
	Role   string `xml:"opf:role,attr,omitempty"`
	FileAs string `xml:"opf:file-as,attr,omitempty"`
	Lang   string `xml:"lang,attr,omitempty"`
	Dir    string `xml:"dir,attr,omitempty"`
}

type Relation struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Coverage struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Rights struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr,omitempty"`
	Lang string `xml:"lang,attr,omitempty"`
	Dir  string `xml:"dir,attr,omitempty"`
}

type Meta struct {
	Text     string `xml:",chardata"`
	Property string `xml:"property,attr,omitempty"`
	Refines  string `xml:"refines,attr,omitempty"`
	ID       string `xml:"id,attr,omitempty"`
	Scheme   string `xml:"scheme,attr,omitempty"`
	Lang     string `xml:"lang,attr,omitempty"`
	Dir      string `xml:"dir,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`
	Content  string `xml:"content,attr,omitempty"`
}

type Link struct {
	Text       string `xml:",chardata"`
	Href       string `xml:"href,attr,omitempty"`
	Rel        string `xml:"rel,attr,omitempty"`
	ID         string `xml:"id,attr,omitempty"`
	Refines    string `xml:"refines,attr,omitempty"`
	MediaType  string `xml:"media-type,attr,omitempty"`
	Properties string `xml:"properties,attr,omitempty"`
}
