package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"TCPChat/functions"
	"TCPChat/models"
)

var list = &models.Users{List: make(map[string]net.Conn)}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on the port %s\n", PORT)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go functions.HandleConnection(conn)
	}
}
