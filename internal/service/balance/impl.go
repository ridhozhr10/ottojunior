package balance

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
)

type balanceServiceImpl struct {
	BalanceRepository repository.Balance
}

// NewService crete implementation for auth.Service
func NewService(
	BalanceRepository repository.Balance,
) Service {
	return &balanceServiceImpl{
		BalanceRepository,
	}
}

func (s *balanceServiceImpl) GetBalance(userID int) (model.Balance, error) {
	res, err := s.BalanceRepository.GetByUserID(userID)
	return res, err
}
