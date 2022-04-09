package dto

type Collection struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
