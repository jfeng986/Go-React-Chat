package router

import (
	"Go-React-Chat/handlers"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()

	// r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/hello", handlers.Hello)
}

func Start(addr string) error {
	return r.Run(addr)
}
