package dto

type LibraryItem struct {
	IsSvg     bool   `json:"isSvg"`
	Cover     []byte `json:"cover"`
	Title     string `json:"title"`
	ItemType  string `json:"itemType"`
	BookCount uint   `json:"bookCount"`
	Id        uint   `json:"id"`
}
