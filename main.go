package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"netcat/functions"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	// Envoyer un message initial au client
	welcomeMessage := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	inputName := "[ENTER YOUR NAME]:"
	userName := ""
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)

	_, err := c.Write([]byte(welcomeMessage))
check:
	_, err2 := c.Write([]byte(inputName))
	if err != nil || err2 != nil {
		fmt.Println("write error:", err, err2)
		return
	}

	for {
		n, err := c.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			fmt.Println("END OF FILE")
			break
		}
		packet = append(packet, tmp[:n]...)
		if userName == "" {
			if !functions.ValidName(packet) {
				goto check
			} else {
				userName = string(packet)
			}
		}

		// Réinitialiser `packet` pour éviter d'envoyer les mêmes données plusieurs fois
		packet = packet[:0]
	}
}