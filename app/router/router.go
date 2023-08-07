package router

import (
	"asterix-golang/app/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SetupRoutes(router *gin.Engine, wsChan chan []byte) {
	// Define the WebSocket route
	router.GET("/geosocket", func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin
		}
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}
		utils.HandleWebSocket(ws, wsChan)
	})
}
