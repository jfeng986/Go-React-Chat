package handlers

import (
	"log"
	"net/http"

	"Go-React-Chat/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	service.Reader(ws)
}
