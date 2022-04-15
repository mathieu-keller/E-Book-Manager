package dto

type Collection struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Books []Book `json:"books"`
}
