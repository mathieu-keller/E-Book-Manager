package dto

import (
	"time"
)

type Book struct {
	ID              uint       `json:"id"`
	Title           string     `json:"title"`
	Published       *time.Time `json:"published"`
	Language        string     `json:"language"`
	Subjects        []Subject  `json:"subjects"`
	Publisher       *string    `json:"publisher"`
	Cover           *[]byte    `json:"cover"`
	Book            string     `json:"book"`
	Authors         []Author   `json:"authors"`
	CollectionId    *uint      `json:"collectionId"`
	CollectionIndex *uint      `json:"collectionIndex"`
}
