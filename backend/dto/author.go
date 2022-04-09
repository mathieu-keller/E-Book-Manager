package dto

type Author struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books"`
}
