package dto

import (
	"time"
)

type Book struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Published    time.Time `json:"published"`
	Language     string    `json:"language"`
	Subject      string    `json:"subject"`
	Publisher    string    `json:"publisher"`
	Cover        []byte    `json:"cover"`
	Book         string    `json:"book"`
	Author       []*Author `json:"author"`
	CollectionId uint
}
