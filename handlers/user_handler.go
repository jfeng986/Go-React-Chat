package handlers

import (
	"log"
	"net/http"

	"Go-React-Chat/models"
	"Go-React-Chat/service"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello"})
}

func Register(c *gin.Context) {
	var userInfo models.User

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
