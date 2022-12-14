package transaction

import (
	"errors"

	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
)

type transactionServiceImpl struct {
	BalanceRepository     repository.Balance
	TransactionRepository repository.Transaction
	UserRepository        repository.User
	ProductRepository     repository.Product
}

// NewService create implementation for transaction service
func NewService(
	BalanceRepository repository.Balance,
	TransactionRepository repository.Transaction,
	UserRepository repository.User,
	ProductRepository repository.Product,
) Service {
	return &transactionServiceImpl{
		BalanceRepository:     BalanceRepository,
		TransactionRepository: TransactionRepository,
		UserRepository:        UserRepository,
		ProductRepository:     ProductRepository,
	}
}

func (s *transactionServiceImpl) GetListTransaction(userID int) ([]model.Transaction, error) {
	data, err := s.TransactionRepository.GetByUserID(userID)
	return data, err
}

func (s *transactionServiceImpl) ConfirmTransaction(userID int, payloads []model.ConfirmTransactionRequest) (model.ConfirmTransactionResponse, error) {
	result := model.ConfirmTransactionResponse{}
	balance, err := s.BalanceRepository.GetByUserID(userID)
	if err != nil {
		return result, err
	}

	for _, payload := range payloads {
		product, err := s.ProductRepository.GetByID(payload.ProductID)
		if err != nil {
			return result, err
		}

		// insert to transaction table
		transaction := model.Transaction{
			Product: product.Product,
			Amount:  uint(payload.Amount),
			Price:   uint64(product.Price),
			Fee:     uint64(product.Fee),
			Total:   uint64(payload.Amount) * (uint64(product.Fee) + uint64(product.Price)),
			UserID:  userID,
		}
		result.Total += int(transaction.Total)
		result.Transaction = append(result.Transaction, transaction)
	}

	// check saldo
	if balance.Total < uint64(result.Total) {
		return result, errors.New("insufficient balance, please do top up to confirm transaction")
	}

	// create transaction
	for _, t := range result.Transaction {
		if _, err := s.TransactionRepository.Create(t); err != nil {
			return result, err
		}
	}

	// update balance
	balance.Total -= uint64(result.Total)
	if _, err := s.BalanceRepository.Update(balance); err != nil {
		return result, err
	}

	result.LatestBalance = int(balance.Total)

	return result, nil
}
