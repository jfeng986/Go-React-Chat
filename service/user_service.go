package service

import (
	"Go-React-Chat/models"
	"Go-React-Chat/repository"
)

func Register(user models.User) (err error) {
	if err = repository.CheckUserExists(user.Username); err != nil {
		return err
	}
	return repository.CreateUser(&user)
}

func Login(user models.User) (err error) {
	return repository.UserAuthentication(user)
}
