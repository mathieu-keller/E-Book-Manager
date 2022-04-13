package dto

type Subject struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
