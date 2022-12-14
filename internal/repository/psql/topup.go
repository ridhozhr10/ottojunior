package psql

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"gorm.io/gorm"
)

type topupPsqlRepository struct {
	DB *gorm.DB
}

// NewTopupPsqlRepository implementation for repository.Topup
func NewTopupPsqlRepository(DB *gorm.DB) repository.Topup {
	return &topupPsqlRepository{DB}
}

func (r *topupPsqlRepository) Create(payload model.Topup) (model.Topup, error) {
	result := r.DB.Create(&payload)
	return payload, result.Error
}
