package repository

import (
	"context"
	"errors"
	"time"

	"Go-React-Chat/db"
	"Go-React-Chat/models"
	"Go-React-Chat/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var collection = db.Client.Database("Go-React-Chat").Collection("users")

func CreateUser(registerRequest *models.RegisterRequest) (*models.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPassword, err := util.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		Username: registerRequest.Username,
		Password: hashedPassword,
		ID:       primitive.NewObjectID().Hex(),
	}
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	registerResponse := &models.RegisterResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
	}
	return registerResponse, nil
}

func CheckUserExists(username string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser := &models.User{}
	err = collection.FindOne(ctx, bson.M{"username": username}).Decode(existingUser)

	switch {
	case err == nil:
		return errors.New("user exists")
	case err != mongo.ErrNoDocuments:
		return err
	default:
		return nil
	}
}

func UserAuthentication(loginRequest models.LoginRequest) (*models.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var storedUser models.User
	err := collection.FindOne(ctx, bson.M{"username": loginRequest.Username}).Decode(&storedUser)
	if err != nil {
		return nil, errors.New("username not found")
	}

	err = util.CheckPassword(storedUser.Password, loginRequest.Password)
	if err != nil {
		return nil, errors.New("password does not match")
	}

	loginResponse := models.LoginResponse{
		ID:       storedUser.ID,
		Username: storedUser.Username,
	}

	return &loginResponse, nil
}

func GetUsers() ([]models.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.GetUsersResponse
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
