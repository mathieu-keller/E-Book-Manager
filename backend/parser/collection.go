package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"strings"
)

/*
<meta property="belongs-to-collection" id="id-2">Keine Cheats für die Liebe</meta>
    <meta refines="#id-2" property="collection-type">series</meta>
*/
func GetCollection(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Meta, cover string) uint {
	var collections = make([]string, 0)
	for _, titleMeta := range *bookFile.Opf.Metadata.Title {
		collection := strings.TrimSpace(titleMeta.Text)
		if metaIdMap["#"+titleMeta.ID]["title-type"].Text == "collection" && collection != "" {
			collections = append(collections, collection)
		}
	}
	for _, metafield := range *bookFile.Opf.Metadata.Meta {
		if strings.HasSuffix(metafield.Name, "series") {
			var collectionName = strings.TrimSpace(metafield.Text)
			if len(collectionName) == 0 {
				collectionName = strings.TrimSpace(metafield.Content)
			}
			if collectionName != "" {
				collections = append(collections, collectionName)
			}
		} else if metafield.Property == "belongs-to-collection" {
			if metaIdMap["#"+metafield.ID]["collection-type"].Text == "series" {
				collections = append(collections, strings.TrimSpace(metafield.Text))
			}
		}
	}
	if len(collections) > 1 {
		return persistCol(collections[0], cover)
	}
	if len(collections) == 0 {
		return 0
	}
	return persistCol(collections[0], cover)
}

func persistCol(title string, cover string) uint {
	var collection = book.GetCollectionByName(title)
	if collection.Title == "" {
		collection.Title = title
		collection.Cover = cover
		collection.Persist()
	} else if collection.Cover == "" {
		collection.Cover = cover
		collection.Persist()
	}
	return collection.ID
}
