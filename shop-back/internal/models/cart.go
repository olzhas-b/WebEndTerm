package models

type Cart struct {
	UserID    uint64 `json:"userID"`
	ProductID uint64 `json:"productID"`
}
