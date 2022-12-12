package model

import "gorm.io/gorm"

// User data structure
type User struct {
	gorm.Model
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// UserRegisterRequest data structure
type UserRegisterRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
