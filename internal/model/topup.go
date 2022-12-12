package model

import "gorm.io/gorm"

// Topup history data structure
type Topup struct {
	gorm.Model
	Amount uint64
	UserID int
	User   User
}
