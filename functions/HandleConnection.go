package functions

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Envoyer un message initial au client
	welcomeMessage := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	inputName := "[ENTER YOUR NAME]:"
	userName := ""
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)

	_, err := conn.Write([]byte(welcomeMessage))
check:
	_, err2 := conn.Write([]byte(inputName))
	if err != nil || err2 != nil {
		fmt.Println("write error:", err, err2)
		return
	}

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			fmt.Println("END OF FILE")
			break
		}
		packet = packet[:0]
		packet = append(packet, tmp[:n]...)
		if userName == "" {
			if !ValidInput(packet) {
				goto check
			}
			userName = strings.Trim(string(packet), "\n")
		}
		now := time.Now()
		message := fmt.Sprintf("[%v][%v]:", now.Format("2006-01-02 15:04:05"), userName)
		_, err3 := conn.Write([]byte(message))
		if err3 != nil {
			fmt.Println("write error:", err3)
			return
		}
	}
}
