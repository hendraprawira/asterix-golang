package utils

import (
	"fmt"
	"net"

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
