package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

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
		start := time.Now().UTC()
		if int(data[0:1][0]) == 240 && len(data) > 400 {
			// value6 := binary.BigEndian.Uint32(data[8:12])
			// fmt.Println(value6)
			datas := AsterixGeoJSONParse(data)
			wsChan <- datas
			processing := time.Since(start)
			fmt.Fprintf(os.Stdout, "\033[0;31m Time taken: %s\033[0m\n ", processing)
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
