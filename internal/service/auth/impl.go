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
	UserRepository    repository.User
	BalanceRepository repository.Balance
}

// NewService crete implementation for auth.Service
func NewService(
	UserRepository repository.User,
	BalanceRepository repository.Balance,
) Service {
	return &authServiceImpl{
		UserRepository,
		BalanceRepository,
	}
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
	balance := model.Balance{
		Total:  0,
		UserID: int(user.ID),
	}
	if _, err := s.BalanceRepository.Create(balance); err != nil {
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
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix() // 10 minute
	claims["data"] = user
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return result, err
	}

	// generate refresh token
	tokenRefresh := jwt.New(jwt.SigningMethodHS256)
	claimsRefresh := tokenRefresh.Claims.(jwt.MapClaims)
	claimsRefresh["exp"] = time.Now().Add(24 * 7 * time.Hour).Unix() // 1 week
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

func (s *authServiceImpl) DecodeToken(tokenRaw string) (int, error) {
	token, err := jwt.Parse(
		tokenRaw,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unauthorized")
			}
			return secretKey, nil
		})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("cant get claims")
	}
	user := claims["data"].(map[string]interface{})
	id := user["ID"].(float64)
	return int(id), nil
}

func (s *authServiceImpl) GetAccountInfo(userID int) (model.User, error) {
	user, err := s.UserRepository.GetByID(userID)
	return user, err
}

func (s *authServiceImpl) hashPass(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func (s *authServiceImpl) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
