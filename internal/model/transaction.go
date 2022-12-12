package model

import "gorm.io/gorm"

// Transaction data structure
type Transaction struct {
	gorm.Model
	Product string
	Amount  uint
	Price   uint64
	Fee     uint64
	Total   uint64
	UserID  int
	User    User
}
