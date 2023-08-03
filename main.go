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
	// port := ":" + os.Getenv("ACTIVE_PORT")

	conn := utils.ConnectionUDP(portUdp)
	packetConn := ipv4.NewPacketConn(conn)
	buffer := make([]byte, 10048)
	defer conn.Close()

	dataChan := make(chan []byte)
	wsChan := make(chan []byte)
	go utils.ReadUDP(packetConn, buffer, dataChan)
	go utils.ProcessData(dataChan, wsChan)

	// Create a new socket.io server
	// server := socketio.NewServer(nil)

	// server.OnConnect("/", func(s socketio.Conn) error {
	// 	s.SetContext("")
	// 	fmt.Println("connected:", s.ID())
	// 	return nil
	// })

	// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	// 	fmt.Println("notice:", msg)
	// 	s.Emit("reply", "have "+msg)
	// })

	// server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
	// 	s.SetContext(msg)
	// 	return "recv " + msg
	// })

	// server.OnEvent("/", "bye", func(s socketio.Conn) string {
	// 	last := s.Context().(string)
	// 	s.Emit("bye", last)
	// 	s.Close()
	// 	return last
	// })

	// server.OnError("/", func(s socketio.Conn, e error) {
	// 	fmt.Println("meet error:", e)
	// })

	// server.OnDisconnect("/", func(s socketio.Conn, reason string) {
	// 	fmt.Println("closed", reason)
	// })

	// go server.Serve()
	// defer server.Close()

	// http.Handle("/socket.io/", server)
	// http.Handle("/", http.FileServer(http.Dir("./asset")))
	// log.Println("Serving at localhost:8000...")
	// log.Fatal(http.ListenAndServe(":8080", nil))

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
	errs := http.ListenAndServe(":8080", nil)
	if errs != nil {
		log.Fatal("Error starting WebSocket server:", errs)
	}
	// // if err := router.Routes().Run(port); err != nil {
	// // 	log.Fatalln(err)
	// // }

	// go utils.HandleWebSockets(server, wsChan)

	select {}
}
