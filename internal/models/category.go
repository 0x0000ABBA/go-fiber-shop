package models

type Category struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ParentCategory *Category `json:"parent_category"`
}
