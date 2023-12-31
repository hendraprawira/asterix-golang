package utils

import (
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"

	"golang.org/x/net/ipv4"
)

func ConnectionUDP(port string) (conn *net.UDPConn) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	conns, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	fmt.Println("UDP Listening Now on Port ", port)
	conn = conns
	return conn
}

func ReadUDP(packetConn *ipv4.PacketConn, buffer []byte, dataChan chan<- []byte) {
	for {
		n, _, _, _ := packetConn.ReadFrom(buffer)
		data := make([]byte, n)
		copy(data, buffer[:n])
		dataChan <- data
	}
}

func ProcessData(dataChan <-chan []byte, wsChan chan<- []byte) {
	for data := range dataChan {
		if int(data[0:1][0]) == 240 && len(data) > 400 {
			datas := AsterixGeoJSONParse(data)
			wsChan <- datas
		}
	}
}

func HandleWebSocket(ws *websocket.Conn, wsChan <-chan []byte) {
	go SendMessages(ws, wsChan)
}

func SendMessages(ws *websocket.Conn, wsChan <-chan []byte) {
	defer ws.Close()
	for data := range wsChan {
		err := ws.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Error sending data over WebSocket:", err)
			break
		}
	}
}
