package dto

type Collection struct {
	ID    uint    `json:"id"`
	Title string  `json:"title"`
	Cover *[]byte `json:"cover"`
	Books []Book  `json:"books"`
}
