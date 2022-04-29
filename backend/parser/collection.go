package parser

import (
	"e-book-manager/book"
	"e-book-manager/epub"
	"gorm.io/gorm"
	"strings"
)

func GetCollection(metaData epub.Metadata, metaIdMap map[string]map[string]epub.Meta, cover *string, tx *gorm.DB) *uint {
	var collections = make([]string, 0)
	if metaData.Title != nil {
		for _, titleMeta := range *metaData.Title {
			collection := strings.TrimSpace(titleMeta.Text)
			if metaIdMap["#"+titleMeta.ID]["title-type"].Text == "collection" && collection != "" {
				collections = append(collections, collection)
			}
		}
	}
	if metaData.Meta != nil {
		for _, metafield := range *metaData.Meta {
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
	}
	if len(collections) > 1 {
		return persistCol(collections[0], cover, tx)
	}
	if len(collections) == 0 {
		return nil
	}
	return persistCol(collections[0], cover, tx)
}

func persistCol(title string, cover *string, tx *gorm.DB) *uint {
	var collection = book.GetCollectionByName(title, tx)
	if collection.Title == "" {
		collection.Title = title
		collection.Cover = cover
		collection.Persist(tx)
	} else if collection.Cover == nil {
		collection.Cover = cover
		collection.Updates(tx)
	}
	return &collection.ID
}
