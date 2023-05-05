package handlers

import (
	"log"
	"net/http"
	"strings"

	"Go-React-Chat/models"
	"Go-React-Chat/service"
	"Go-React-Chat/util"

	"github.com/gin-gonic/gin"
)

var userInfo models.User

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}

func Register(c *gin.Context) {
	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to BindJSON"})
		return
	}

	if err := service.Register(userInfo); err != nil {
		log.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Registered"})
}

func Login(c *gin.Context) {
	if err := c.BindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to BindJSON"})
		return
	}

	if err := service.Login(userInfo); err != nil {
		log.Printf("Failed to login: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := util.GenerateToken(userInfo.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged In"})
}

func JwtAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	log.Println("tokenString: ", tokenString)
	if tokenString == "" {
		// log.Println("No token provided")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}

	token, err := util.ParseToken(tokenString)
	if err != nil || token == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Valid token"})
	}
}
