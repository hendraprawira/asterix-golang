package main

import (
	"asterix-golang/app/router"
	"asterix-golang/app/utils"
	"fmt"
	"log"
	"os"

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
	recHost := os.Getenv("HOST_UDP_REC")

	conn := utils.ConnectionUDP(portUdp)
	packetConn := ipv4.NewPacketConn(conn)
	buffer := make([]byte, 10048)
	defer conn.Close()

	go utils.ReadUDP(packetConn, buffer, recHost)

	if err := router.Routes().Run(port); err != nil {
		log.Fatalln(err)
	}

	select {}
}
