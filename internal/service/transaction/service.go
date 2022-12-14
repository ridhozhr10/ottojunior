package transaction

import "github.com/ridhozhr10/ottojunior/internal/model"

// Service definition
type Service interface {
	GetListTransaction(userID int) ([]model.Transaction, error)
	ConfirmTransaction(int, []model.ConfirmTransactionRequest) (model.ConfirmTransactionResponse, error)
}
