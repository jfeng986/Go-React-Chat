package handlers

import (
	"log"
	"net/http"
	"sync"

	"Go-React-Chat/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	onlineUsers = make(map[string]bool)
	mutex       sync.Mutex
)

func WsHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	userID := r.URL.Query().Get("userID")
	mutex.Lock()
	onlineUsers[userID] = true
	mutex.Unlock()

	go func() {
		defer ws.Close()
		for {
			// Read messages from the client
			_, _, err := ws.ReadMessage()
			if err != nil {
				break
			}
		}

		// When the client disconnects, remove them from the online users map
		mutex.Lock()
		delete(onlineUsers, userID)
		mutex.Unlock()
	}()

	service.Reader(ws)
}
