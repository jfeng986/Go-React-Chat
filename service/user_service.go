package service

import (
	"fmt"
	"log"

	"Go-React-Chat/models"
	"Go-React-Chat/repository"

	"github.com/gorilla/websocket"
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
