package topup

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service definition for topup
type Service interface {
	TopupBalance(model.TopupBalanceRequest) (model.TopupBalanceResponse, error)
}
