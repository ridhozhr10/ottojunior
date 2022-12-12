package model

// Product data structure
type Product struct {
	ID          int    `json:"id"`
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Fee         int    `json:"fee"`
}
