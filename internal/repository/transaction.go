package repository

import "github.com/ridhozhr10/ottojunior/internal/model"

// Transaction repository definition
type Transaction interface {
	Create(model.Transaction) (model.Transaction, error)
	GetByID(int) ([]model.Transaction, error)
}
