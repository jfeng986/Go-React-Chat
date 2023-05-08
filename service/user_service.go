package service

import (
	"Go-React-Chat/models"
	"Go-React-Chat/repository"
)

func Register(registerRequest models.RegisterRequest) (*models.RegisterResponse, error) {
	if err := repository.CheckUserExists(registerRequest.Username); err != nil {
		return nil, err
	}
	return repository.CreateUser(&registerRequest)
}

func Login(loginRequest models.LoginRequest) (*models.LoginResponse, error) {
	loginResponse, err := repository.UserAuthentication(loginRequest)
	if err != nil {
		return nil, err
	}
	return loginResponse, nil
}

func GetUsers() ([]models.GetUsersResponse, error) {
	users, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
