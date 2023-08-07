package main

import (
	"asterix-golang/app/router"
	"asterix-golang/app/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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
	defer conn.Close()

	buffer := make([]byte, 10048)
	dataChan := make(chan []byte)
	wsChan := make(chan []byte)

	go utils.ReadUDP(packetConn, buffer, dataChan)
	go utils.ProcessData(dataChan, wsChan)

	routers := gin.Default()
	router.SetupRoutes(routers, wsChan)

	errs := routers.Run(port)
	if errs != nil {
		log.Fatal("Error starting WebSocket server:", errs)
	}

	select {}
}
