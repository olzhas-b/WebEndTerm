package models

type Product struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	CategoryID uint64  `json:"categoryID"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
}

type ProductDTO struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	CategoryName string  `json:"categoryName"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	IsFavorite   bool    `json:"isFavorite"`
}
