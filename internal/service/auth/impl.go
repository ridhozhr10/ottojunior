package auth

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	UserRepository repository.User
}

// NewService crete implementation for auth.Service
func NewService(
	UserRepository repository.User,
) Service {
	return &authServiceImpl{UserRepository}
}

func (s *authServiceImpl) Register(payload model.User) (model.User, error) {
	pass, err := s.hashPass(payload.Password)
	if err != nil {
		return payload, err
	}
	payload.Password = pass
	user, err := s.UserRepository.Create(payload)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *authServiceImpl) hashPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
