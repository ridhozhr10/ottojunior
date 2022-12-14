package product

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service auth definition
type Service interface {
	GetProductList() ([]model.Product, error)
}
