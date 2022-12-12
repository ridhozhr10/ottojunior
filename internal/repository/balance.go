package repository

import "github.com/ridhozhr10/ottojunior/internal/model"

// Balance repository definition
type Balance interface {
	GetByUserID(int) (model.Balance, error)
	Update(model.Balance) (model.Balance, error)
}
