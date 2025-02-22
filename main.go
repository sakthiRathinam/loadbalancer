package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	startTCPServer()
}

// type server struct {
// 	address        string
// 	headers        map[string]interface{}
// 	healthCheckURL string
// 	active         bool
// }
// type loadbalancer struct {
// 	servers []string
// }

func startTCPServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer))
	}
}
