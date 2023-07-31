package main

import (
	"asterix-golang/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	websocketGo "asterix-golang/utils/websocket"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"golang.org/x/net/ipv4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("UDP_PORT")
	recHost := os.Getenv("HOST_UDP_REC")
	conn := utils.ConnectionUDP(port)
	packetConn := ipv4.NewPacketConn(conn)
	buffer := make([]byte, 10048)
	defer conn.Close()

	go utils.ReadUDP(packetConn, buffer, recHost)

	http.HandleFunc("/geosocket", WebSocket)
	fmt.Println("WebSocket server is running on http://localhost:8000/ws")
	http.ListenAndServe(":8080", nil)
	select {}
}

func WebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketGo.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[BE-Tracking-Surface-Debug] [err] [controllers-tracking] [WebSocketGetDataSurface] [Upgrade Connection]: ", err)
		return
	}

	defer conn.Close()
	websocketGo.GlobalWebSocketCon = conn // Store the connection globally

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
