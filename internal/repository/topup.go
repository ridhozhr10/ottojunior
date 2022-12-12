package repository

import "github.com/ridhozhr10/ottojunior/internal/model"

// Topup repository definition
type Topup interface {
	Create(model.Topup) (model.Topup, error)
}
