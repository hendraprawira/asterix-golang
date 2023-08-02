package utils

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/ipv4"
)

func ConnectionUDP(port string) (conn *net.UDPConn) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	// Create a UDP connection
	conns, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	fmt.Println("UDP Listening Now on Port ", port)

	conn = conns
	return conn
}

func ReadUDP(packetConn *ipv4.PacketConn, buffer []byte, host string) {
	for {
		start := time.Now().UTC()
		n, _, _, _ := packetConn.ReadFrom(buffer)
		data := buffer[:n]
		if int(data[0:1][0]) == 240 && n > 500 {
			AsterixGeoJSONParse(data)
			processing := time.Since(start)
			fmt.Fprintf(os.Stdout, "\033[0;31m Time taken: %s\033[0m\n ", processing)
		}
	}

}

func SendDataUdp(host string, data []byte) {
	destAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println("Error resolving destination address:", err)
		return
	}

	// Create UDP connection to send data
	sendConn, err := net.DialUDP("udp", nil, destAddr)
	if err != nil {
		fmt.Println("Error creating UDP connection for sending:", err)
		return
	}
	defer sendConn.Close()

	// Forward the data to the destination
	_, err = sendConn.Write(data)
	if err != nil {
		fmt.Println("Error forwarding data:", err)
		return
	}
}
