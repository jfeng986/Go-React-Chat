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

type jwtAuthResponse struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func Register(c *gin.Context) {
	var registerRequest models.RegisterRequest
	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to BindJSON"})
		return
	}

	registerResponse, err := service.Register(registerRequest)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Registered", "registerResponse": registerResponse})
}

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to BindJSON"})
		return
	}

	loginResponse, err := service.Login(loginRequest)
	if err != nil {
		log.Printf("Failed to login: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := util.GenerateToken(loginResponse.Username, loginResponse.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged In", "loginResponse": loginResponse})
}

func JwtAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}

	token, err := util.ParseToken(tokenString)
	if err != nil || token == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
	} else {
		jwtAuthResponse := jwtAuthResponse{
			Username: token.Username,
			ID:       token.ID,
		}
		c.JSON(http.StatusOK, gin.H{"message": "Valid token", "jwtAuthResponse": jwtAuthResponse})
	}
}
