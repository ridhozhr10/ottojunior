package psql

import (
	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
	"gorm.io/gorm"
)

type userPsqlRepository struct {
	DB *gorm.DB
}

// NewUserPsqlRepository create new repository implementation using gorm db
func NewUserPsqlRepository(DB *gorm.DB) repository.User {
	return &userPsqlRepository{DB}
}

func (p *userPsqlRepository) GetByUsernameEmail(param string) (model.User, error) {
	user := model.User{}
	err := p.DB.Where("username = ?", param).Or("email = ?", param).First(&user).Error
	return user, err
}

func (p *userPsqlRepository) Create(data model.User) (model.User, error) {
	err := p.DB.Create(&data).Error
	return data, err
}

func (p *userPsqlRepository) GetByID(id int) (model.User, error) {
	user := model.User{}
	err := p.DB.First(&user, id).Error
	return user, err
}

func (p *userPsqlRepository) GetByPhoneNumber(phoneNumber string) (model.User, error) {
	user := model.User{}
	err := p.DB.Where("phone_number = ?", phoneNumber).First(&user).Error
	return user, err
}
