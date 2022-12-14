package product

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
)

type productServiceImpl struct {
	ProductRepository repository.Product
}

// NewService crete implementation for auth.Service
func NewService(
	ProductRepository repository.Product,
) Service {
	return &productServiceImpl{
		ProductRepository,
	}
}

func (s *productServiceImpl) GetProductList() ([]model.Product, error) {
	res, err := s.ProductRepository.Get()
	return res, err
}
