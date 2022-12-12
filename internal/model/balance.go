package model

import "gorm.io/gorm"

// Balance latest data structure
type Balance struct {
	gorm.Model
	Total  uint64
	UserID int
	User   User
}
