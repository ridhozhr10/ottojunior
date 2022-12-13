package auth

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service auth definition
type Service interface {
	Register(model.User) (model.User, error)
	Login(model.UserLoginRequest) (model.UserLoginResponse, error)
	GetAccountInfo(int) (model.User, error)
	DecodeToken(token string) (int, error)
}
