package model

import "gorm.io/gorm"

// Balance latest data structure
type Balance struct {
	gorm.Model
	Total  uint64 `json:"total"`
	UserID int    `json:"user_id"`
	User   User   `json:"-"`
}
