package router

import (
	"log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin
}

// Define a structure to represent connected clients.
type Client struct {
	conn *websocket.Conn
}

var clients = make(map[*Client]bool)

func SetupRoutes(router *gin.Engine, wsChan chan []byte) {
	// Define the WebSocket route
	router.GET("/geosocket", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}

		defer ws.Close()

		client := &Client{conn: ws}
		clients[client] = true

		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
		}

		delete(clients, client)
	})

	go HandleWebSocket(clients, wsChan)
}

func HandleWebSocket(clients map[*Client]bool, wsChan <-chan []byte) {
	for data := range wsChan {
		for client := range clients {
			err := client.conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Error sending data over WebSocket:", err)
				client.conn.Close()
				delete(clients, client)
				break
			}
		}
	}
}
