package main

import (
	"asterix-golang/app/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"golang.org/x/net/ipv4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	portUdp := os.Getenv("UDP_PORT")
	port := ":" + os.Getenv("ACTIVE_PORT")

	conn := utils.ConnectionUDP(portUdp)
	packetConn := ipv4.NewPacketConn(conn)
	buffer := make([]byte, 10048)
	defer conn.Close()

	dataChan := make(chan []byte)
	wsChan := make(chan []byte)

	go utils.ReadUDP(packetConn, buffer, dataChan)
	go utils.ProcessData(dataChan, wsChan)

	http.HandleFunc("/geosocket", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin
		}
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}
		utils.HandleWebSocket(ws, wsChan)
	})

	// Start the WebSocket server
	errs := http.ListenAndServe(port, nil)
	if errs != nil {
		log.Fatal("Error starting WebSocket server:", errs)
	}

	select {}
}
