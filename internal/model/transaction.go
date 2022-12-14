package model

import "gorm.io/gorm"

// Transaction data structure
type Transaction struct {
	gorm.Model
	Product string `json:"product"`
	Amount  uint   `json:"amount"`
	Price   uint64 `json:"price"`
	Fee     uint64 `json:"fee"`
	Total   uint64 `json:"total"`
	UserID  int    `json:"user_id"`
	User    User   `json:"-"`
}

// ConfirmTransactionRequest data structure
type ConfirmTransactionRequest struct {
	ProductID int `json:"product_id"`
	Amount    int `json:"amount"`
}

// ConfirmTransactionResponse data structure
type ConfirmTransactionResponse struct {
	Total         int           `json:"total"`
	LatestBalance int           `json:"latest_balance"`
	Transaction   []Transaction `json:"transaction"`
}
