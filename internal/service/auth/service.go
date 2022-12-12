package auth

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service auth definition
type Service interface {
	Register(model.User) (model.User, error)
}
