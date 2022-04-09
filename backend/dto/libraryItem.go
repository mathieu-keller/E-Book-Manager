package dto

type LibraryItem struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Cover []byte `json:"cover"`
}
