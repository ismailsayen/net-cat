package functions

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	welcomeMessage := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	userName := ""

	chanErr := make(chan error, 2)
	_, err := conn.Write([]byte(welcomeMessage))
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	chanErr <- handleName(conn, &userName)

	chanErr <- handleMessage(conn, userName)

	for er := range chanErr {
		if er != nil {
			return
		}
	}
}
