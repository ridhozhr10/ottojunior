package repository

import "github.com/ridhozhr10/ottojunior/internal/model"

// User repository defintion
type User interface {
	GetByUsernameEmail(string) (model.User, error)
	Create(model.User) (model.User, error)
	GetByID(int) (model.User, error)
	GetByPhoneNumber(string) (model.User, error)
}
