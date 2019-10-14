// tcp_server example in go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_PORT = ":8080"
	CONN_TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received from the client: ", string(message))
	conn.Write([]byte(message + "\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_PORT)

	if err != nil {
		log.Fatal("Error starting tcp server: ", err)
	}

	defer listener.Close()

	log.Println("Listening on: " + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
