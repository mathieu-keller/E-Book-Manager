package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"strings"
)

func GetCollection(bookFile *epub.Book, metaIdMap map[string]map[string]epub.Metafield) uint {
	var collections = make([]string, 0)
	for _, titleMeta := range bookFile.Opf.Metadata.Title {
		collection := strings.TrimSpace(titleMeta.Data)
		if metaIdMap["#"+titleMeta.ID]["title-type"].Data == "collection" && collection != "" {
			collections = append(collections, collection)
		}
	}
	for _, metafield := range bookFile.Opf.Metadata.Meta {
		if strings.HasSuffix(metafield.Name, "series") {
			var collectionName = strings.TrimSpace(metafield.Data)
			if len(collectionName) == 0 {
				collectionName = strings.TrimSpace(metafield.Content)
			}
			if collectionName != "" {
				collections = append(collections, collectionName)
			}
		}
	}
	if len(collections) > 1 {
		return persistCol(collections[0])
	}
	if len(collections) == 0 {
		return 0
	}
	return persistCol(collections[0])
}

func persistCol(title string) uint {
	var collection = book.GetCollectionByName(title)
	if collection.Title == "" {
		collection.Title = title
		collection.Persist()
	}
	return collection.ID
}

/*
if bookEntity.CollectionId == 0 {
		for _, metafield := range bookFile.Opf.Metadata.Meta {
			if strings.HasSuffix(metafield.Name, "series") {
				var collectionName = strings.TrimSpace(metafield.Data)
				if len(collectionName) == 0 {
					collectionName = strings.TrimSpace(metafield.Content)
				}
				var collection = book.GetCollectionByName(collectionName)
				if collection.Name == "" {
					collection.Name = collectionName
					collection.Persist()
				}
				bookEntity.CollectionId = collection.ID
				break
			}
		}
	}
*/
