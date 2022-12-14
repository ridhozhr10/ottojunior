package model

import "gorm.io/gorm"

// Topup history data structure
type Topup struct {
	gorm.Model
	Amount uint64
	UserID int
	User   User
}

// TopupBalanceRequest data structure
type TopupBalanceRequest struct {
	Amount      int    `json:"amount"`
	PhoneNumber string `json:"phone_number"`
}

// TopupBalanceResponse data structure
type TopupBalanceResponse struct {
	LatestBalance int `json:"latest_balance"`
}
