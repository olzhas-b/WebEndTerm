package models

type Product struct {
	ID         uint64  `json:"id,omitempty"`
	Name       string  `json:"name"`
	CategoryID uint64  `json:"categoryID"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
}
