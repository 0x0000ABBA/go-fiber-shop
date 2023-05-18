package models

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Category    Category `json:"category"`
	Prices      []Price  `json:"prices"`
	Images      []Image  `json:"images"`
	Reviews     []Review `json:"reviews"`
}
