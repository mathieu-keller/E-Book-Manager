package dto

type Subject struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
