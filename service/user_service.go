package service

import (
	"fmt"
	"log"

	"Go-React-Chat/models"
	"Go-React-Chat/repository"

	"github.com/gorilla/websocket"
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

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
