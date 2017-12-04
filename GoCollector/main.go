package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// We will be listening for Metrics over a UDP port exposed to localhost. This will make it
	// easy to pubish data to the collector without the overhead of TCP.
	CollectorAddr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		log.Fatal(err)
	}

	// Create the server
	CollectorConn, err := net.ListenUDP("udp", CollectorAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer CollectorConn.Close()

	// Creating our UDP byte buffer
	buf := make([]byte, 1024)

	for {
		i, _, err := CollectorConn.ReadFromUDP(buf)
		fmt.Println("Received: ", string(buf[0:i]))

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
