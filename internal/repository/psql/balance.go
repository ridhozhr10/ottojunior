package psql

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"gorm.io/gorm"
)

type balancePsqlRepository struct {
	DB *gorm.DB
}

// NewBalancePsqlRepository return implementation for repository.Balance
func NewBalancePsqlRepository(DB *gorm.DB) repository.Balance {
	return &balancePsqlRepository{DB}
}

func (r *balancePsqlRepository) Create(payload model.Balance) (model.Balance, error) {
	result := r.DB.Create(&payload)
	return payload, result.Error
}

func (r *balancePsqlRepository) GetByUserID(userID int) (model.Balance, error) {
	result := model.Balance{}
	q := r.DB.Where("user_id = ?", userID).First(&result)
	return result, q.Error
}

func (r *balancePsqlRepository) Update(payload model.Balance) (model.Balance, error) {
	q := r.DB.Save(&payload)
	return payload, q.Error
}
