package asterix

import (
	"asterix-golang/app/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebSocket(c *gin.Context) {
	conn, err := utils.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("[BE-Tracking-Surface-Debug] [err] [controllers-tracking] [WebSocketGetDataSurface] [Upgrade Connection]: ", err)
		return
	}

	defer conn.Close()
	utils.GlobalWebSocketCon = conn // Store the connection globally

	for {
		// Read message from WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("[BE-Tracking-Surface-Debug] [err] [controllers-tracking] [WebSocketGetDataSurface] [Read Message]: ", err)
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("WebSocket connection closed.")
			} else {
				log.Printf("Error reading WebSocket message: %v", err)
			}
			break
		}
		log.Println("Received message from WebSocket:", string(message))
		response := []byte("Received your message: " + string(message))
		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			log.Println("[BE-Tracking-Surface-Debug] [err] [controllers-tracking] [WebSocketGetDataSurface] [Send Responese]: ", err)
			break
		}
	}
}
