package topup

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
)

type topupServiceImpl struct {
	UserRepository    repository.User
	TopupRepository   repository.Topup
	BalanceRepository repository.Balance
}

// NewService crete implementation for topup.Service
func NewService(
	UserRepository repository.User,
	TopupRepository repository.Topup,
	BalanceRepository repository.Balance,
) Service {
	return &topupServiceImpl{
		UserRepository,
		TopupRepository,
		BalanceRepository,
	}
}

func (s *topupServiceImpl) TopupBalance(payload model.TopupBalanceRequest) (model.TopupBalanceResponse, error) {
	result := model.TopupBalanceResponse{}
	user, err := s.UserRepository.GetByPhoneNumber(payload.PhoneNumber)
	if err != nil {
		return result, err
	}
	balance, err := s.BalanceRepository.GetByUserID(int(user.ID))
	if err != nil {
		return result, err
	}
	balance.Total += uint64(payload.Amount)
	if _, err := s.BalanceRepository.Update(balance); err != nil {
		return result, err
	}

	// create topup
	topup := model.Topup{
		Amount: uint64(payload.Amount),
		UserID: int(user.ID),
	}
	if _, err := s.TopupRepository.Create(topup); err != nil {
		return result, err
	}
	result.LatestBalance = int(balance.Total)
	return result, nil
}
