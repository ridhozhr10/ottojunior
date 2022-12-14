package balance

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service auth definition
type Service interface {
	GetBalance(int) (model.Balance, error)
}
