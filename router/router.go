package router

import (
	"time"

	"Go-React-Chat/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/jwtauth", handlers.JwtAuth)
	r.GET("/ws", handlers.WsHandler)
	// r.GET("/profile", handlers.Profile)
}

func Start(addr string) error {
	return r.Run(addr)
}
