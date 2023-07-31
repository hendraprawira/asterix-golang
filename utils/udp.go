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
		n, _, _, err := packetConn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		data := buffer[:n]
		if int(data[0:1][0]) == 240 && n > 500 {
			AsterixGeoJSONParse(data)
			// jsonData, _ := json.Marshal(dataStruct.I041.StartAz)
			// fmt.Println(dataStruct.I041.StartAz)
			// SendDataUdp(host, jsonData)
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
