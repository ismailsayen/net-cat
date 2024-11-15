package functions

import (
	"fmt"
	"net"
	"os"
	"sync"

	"TCPChat/models"
)

var (
	list      = &models.Users{List: make(map[string]net.Conn)}
	mu        sync.Mutex
	UsersConn int
	messages  [][]byte
)

func HandleConnection(conn net.Conn) {
	UsersConn++
	defer conn.Close()
	welcomeMessage := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	userName := ""
	if UsersConn > 3 {
		conn.Write([]byte("The list has exceeded the maximum allowed size of 10 items."))
		os.Exit(1)
	}
	mu.Lock()
	_, err := conn.Write([]byte(welcomeMessage))
	mu.Unlock()
	if err != nil {
		fmt.Println("write error:", err)
		return
	}
	for i := range messages {
		if string(messages[i]) == "\n" {
			continue
		}
		fmt.Println(string(messages[i]))
	}
again:
	Err := handleName(conn, &userName)
	for t := range list.List {
		if t == userName || Err != nil {
			mu.Lock()
			conn.Write([]byte("This name: " + userName + "is already taken, please choose another name.\n"))
			mu.Unlock()
			userName = ""
			goto again
		}
	}

	list.List[userName] = conn
	joiningMsg := fmt.Sprintf("\033[2K\r%s has joined our chat...\n", userName)
	mu.Lock()
	errB := BrodcastMsg(joiningMsg, conn)
	mu.Unlock()
	if errB != nil {
		return
	}
	chanErr2 := handleMessage(conn, userName)
	if chanErr2 != nil {
		left_msg := fmt.Sprintf("\033[2K\r%s has left our chat...\n", userName)
		errB := BrodcastMsg(left_msg, conn)
		if errB != nil {
			return
		}
		delete(list.List, userName)
		UsersConn--
	}
}
