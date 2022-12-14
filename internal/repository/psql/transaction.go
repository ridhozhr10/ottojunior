package psql

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"gorm.io/gorm"
)

type transactionPsqlRepository struct {
	DB *gorm.DB
}

// NewTransactionPsqlRepository implementation for repository.Transaction
func NewTransactionPsqlRepository(DB *gorm.DB) repository.Transaction {
	return &transactionPsqlRepository{DB}
}

func (r *transactionPsqlRepository) Create(payload model.Transaction) (model.Transaction, error) {
	result := r.DB.Create(&payload)
	return payload, result.Error
}

func (r *transactionPsqlRepository) GetByID(userID int) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}
