package repository

import (
	"gorm.io/gorm"
	"technopartner/internal/entity"
)

type AuthRepo interface {
	GetUserByEmail(email string) (entity.User, error)
	Login(email string) (entity.User, error)
	Register(user entity.RegisterRequest) (entity.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func (a authRepo) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (a authRepo) Login(email string) (entity.User, error) {
	var user entity.User
	result := a.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (a authRepo) Register(request entity.RegisterRequest) (entity.User, error) {
	user := entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		RoleId:   "56cb7a92-580a-4b18-82f5-93f30a4231a4",
	}

	result := a.db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func New(db *gorm.DB) *authRepo {
	return &authRepo{
		db,
	}
}
