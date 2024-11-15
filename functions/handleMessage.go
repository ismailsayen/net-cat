package functions

import (
	"fmt"
	"net"
	"time"

	"TCPChat/utils"
)

func handleMessage(conn net.Conn, userName string) error {
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)
WriteAgain:
	now := time.Now()
	message := fmt.Sprintf("[%v][%v]:", now.Format(time.DateTime), userName)
	utils.Writer(message, conn)
	n, err2 := conn.Read(tmp)

	if err2 != nil {
		return err2
	}

	packet = packet[:0]
	packet = append(packet, tmp[:n]...)

	if !ValidInput(packet) {
		goto WriteAgain
	}

	message2 := fmt.Sprintf("\n[%v][%v]:%v", now.Format("2006-01-02 15:04:05"), userName, string(packet))

	errB := BrodcastMsg(message2, conn)

	if errB != nil {
		return errB
	}
	messages = append(messages, packet)
	goto WriteAgain
}
