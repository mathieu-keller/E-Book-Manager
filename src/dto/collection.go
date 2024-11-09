package dto

type Collection struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Cover *string `json:"cover"`
	Books []Book  `json:"books"`
}
