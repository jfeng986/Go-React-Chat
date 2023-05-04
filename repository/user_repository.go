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
	"go.mongodb.org/mongo-driver/mongo" // for BSON IDs
)

var collection = db.Client.Database("Go-React-Chat").Collection("users")

func CreateUser(user *models.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if user.Password, err = util.HashPassword(user.Password); err != nil {
		return err
	}

	user.ID = primitive.NewObjectID().Hex()
	_, err = collection.InsertOne(ctx, user)
	return err
}

func CheckUserExists(username string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser := &models.User{}
	err = collection.FindOne(ctx, bson.M{"username": username}).Decode(existingUser)

	switch {
	case err == nil:
		return errors.New("username already exists")
	case err != mongo.ErrNoDocuments:
		return err
	default:
		return nil
	}
}
