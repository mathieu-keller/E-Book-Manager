package dto

type LibraryItem struct {
	Cover     []byte `json:"cover"`
	Name      string `json:"name"`
	ItemType  string `json:"itemType"`
	BookCount uint   `json:"bookCount"`
	Id        uint   `json:"id"`
}
