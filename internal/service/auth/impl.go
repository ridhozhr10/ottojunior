package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key")

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

func (s *authServiceImpl) Login(payload model.UserLoginRequest) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}
	user, err := s.UserRepository.GetByUsernameEmail(payload.UserIdentity)
	if err != nil {
		return result, err
	}
	if !s.checkPasswordHash(payload.Password, user.Password) {
		return result, errors.New("incorrect username or password")
	}

	// generate access token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute) // 10 minute
	claims["data"] = user
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return result, err
	}

	// generate refresh token
	tokenRefresh := jwt.New(jwt.SigningMethodHS256)
	claimsRefresh := tokenRefresh.Claims.(jwt.MapClaims)
	claimsRefresh["exp"] = time.Now().Add(24 * 7 * time.Hour) // 1 week
	claimsRefresh["data"] = user
	refreshToken, err := tokenRefresh.SignedString(secretKey)
	if err != nil {
		return result, err
	}

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	result.Data = user

	return result, nil
}

func (s *authServiceImpl) hashPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func (s *authServiceImpl) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
