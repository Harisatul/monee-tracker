package service

import (
	"errors"
	"technopartner/internal/entity"
	"technopartner/internal/repository"
	"technopartner/pkg/hash"
	"technopartner/pkg/validator"
)

type AuthService interface {
	Login(request entity.LoginRequest) (entity.LoginResponse, error)
	Register(user entity.RegisterRequest) (entity.RegisterResponse, error)
}

type authService struct {
	authRepo repository.AuthRepo
}

func New(adminRepo repository.AuthRepo) *authService {
	return &authService{
		adminRepo,
	}
}

func (ac *authService) Register(request entity.RegisterRequest) (entity.RegisterResponse, error) {
	if err := validator.Validation(request); err != nil {
		return entity.RegisterResponse{}, err
	}

	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		return entity.RegisterResponse{}, err
	}

	request.Password = string(hashedPassword)

	register, err := ac.authRepo.Register(request)
	if err != nil {
		return entity.RegisterResponse{}, err
	}

	response := entity.RegisterResponse{
		Email:    register.Email,
		Username: register.Username,
		Password: register.Password,
	}

	return response, nil
}

func (ac *authService) Login(request entity.LoginRequest) (entity.LoginResponse, error) {
	if err := validator.Validation(request); err != nil {
		return entity.LoginResponse{}, err
	}

	login, err := ac.authRepo.Login(request.Email)

	if err != nil {
		return entity.LoginResponse{}, errors.New("Email atau password salah")

	}

	err = hash.VerifyPassword(login.Password, request.Password)
	if err != nil {
		return entity.LoginResponse{}, errors.New("Email atau password salah")
	}

	//
	//token, err := jwt.CreateToken(response.ID, response.Email)
	//if err != nil {
	//	return nil, "", err
	//}
	//response.Token = token

	response := entity.LoginResponse{
		Email: login.Email,
	}

	return response, nil
}
