package models

type Favorite struct {
	UserID    uint64 `json:"userID"`
	ProductID uint64 `json:"productID"`
}
