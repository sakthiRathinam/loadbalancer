package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

func main() {
	startTCPServer()
}

type server struct {
	address        string
	headers        map[string]interface{}
	healthCheckURL string
	active         bool
}

func startTCPServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	servers := []server{
		{
			address:        "10.0.0.2:8000",
			healthCheckURL: "http://10.0.0.2:8000/health",
			active:         true,
			headers:        map[string]interface{}{},
		},
		{
			address:        "10.0.0.3:8000",
			healthCheckURL: "http://10.0.0.3:8000/health",
			active:         true,
			headers:        map[string]interface{}{},
		},
		{
			address:        "10.0.0.5:8000",
			healthCheckURL: "http://10.0.0.5:8000/health",
			active:         true,
			headers:        map[string]interface{}{},
		},
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		server := pickRandomServer(servers)
		go handleConnection(conn, server)
	}
}

func pickRandomServer(servers []server) server {
	return servers[rand.Intn(len(servers))]
}

func handleConnection(conn net.Conn, server server) {
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
