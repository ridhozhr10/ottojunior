package repository

import "github.com/ridhozhr10/ottojunior/internal/model"

// Product repository definition
type Product interface {
	Get() ([]model.Product, error)
	GetByID(int) (model.Product, error)
}
