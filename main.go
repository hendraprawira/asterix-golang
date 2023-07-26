package main

import (
	"asterix-golang/utils"
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

func main() {
	// Define the UDP address to listen on
	// address := "172.16.21.205:4378"
	// address2 := "172.16.6.168:8003"

	addr, err := net.ResolveUDPAddr("udp", ":"+"8003")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()
	packetConn := ipv4.NewPacketConn(conn)
	// Create a buffer to hold incoming data
	buffer := make([]byte, 6048)
	go utils.StartUDP(packetConn, buffer)
	select {}

}
