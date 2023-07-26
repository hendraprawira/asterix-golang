package utils

import (
	"encoding/json"
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

// func StartUDP(packetConn *ipv4.PacketConn, buffer []byte) {
// 	for {
// 		// Read from the connection
// 		// start := time.Now().UTC()

// 		n, _, _, err := packetConn.ReadFrom(buffer)
// 		if err != nil {
// 			fmt.Println("Error reading:", err)
// 			continue
// 		}

// 		// Process the received data
// 		data := buffer[:n]
// 		hexString := hex.EncodeToString(data)
// 		var result []string

// 		for i := 0; i < len(hexString); i += 2 {
// 			if i+2 <= len(hexString) {
// 				result = append(result, hexString[i:i+2])
// 			} else {
// 				result = append(result, hexString[i:])
// 			}
// 		}
// 		if result[0] == "f0" && len(result) > 510 {
// 			dataStruct := AsterixParse(result)
// 			// jsonData, _ := json.Marshal(dataStruct)
// 			// SendDataUdp("172.16.6.168:8003", jsonData)
// 			// Print the JSON data
// 			fmt.Println(dataStruct)
// 			// fmt.Println(string(jsonData))
// 			// processing := time.Since(start)
// 			// fmt.Fprintf(os.Stdout, "\033[0;31m Time taken: %s\033[0m\n ", processing
// 		}

// 	}

// }

func StartUDP(packetConn *ipv4.PacketConn, buffer []byte) {
	expect := 938906
	for {
		n, _, _, err := packetConn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		data := buffer[:n]
		if int(data[0:1][0]) == 240 && len(data) > 510 {
			dataStruct := AsterixParse(data)
			if int(dataStruct.I020.MsgIndex) != expect {
				fmt.Println("Looping tidak berurutan. Looping dihentikan.")
				break
			}
			expect = int(dataStruct.I020.MsgIndex) + 1
			jsonData, _ := json.Marshal(dataStruct)
			fmt.Println(string(jsonData))
			// log.Println(dataStruct)
			SendDataUdp("172.16.6.168:8003", jsonData)

		}

	}

}

// func StartUDP(packetConn *ipv4.PacketConn, buffer []byte) {
// 	expect := 938906
// 	for {
// 		n, _, _, err := packetConn.ReadFrom(buffer)
// 		if err != nil {
// 			fmt.Println("Error reading:", err)
// 			continue
// 		}
// 		data := buffer[:n]
// 		if int(data[0:1][0]) == 240 && len(data) > 510 {
// 			dataStruct := binary.BigEndian.Uint32(data[8:12])

// 			if int(dataStruct) != expect {
// 				fmt.Println("Looping tidak berurutan. Looping dihentikan.")
// 				break
// 			}
// 			expect = int(dataStruct) + 1
// 			fmt.Println(dataStruct)
// 		}

// 	}

// }

func SendDataUdp(destinationPort string, data []byte) {
	destAddr, err := net.ResolveUDPAddr("udp", destinationPort)
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
