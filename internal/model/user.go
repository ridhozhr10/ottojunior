package model

import "gorm.io/gorm"

// User data structure
type User struct {
	gorm.Model
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
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

// UserLoginRequest data structure
type UserLoginRequest struct {
	UserIdentity string `json:"user_identity"`
	Password     string `json:"password"`
}

// UserLoginResponse data structure
type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Data         User   `json:"data"`
}
