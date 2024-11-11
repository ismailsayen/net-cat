package functions

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func handleMessage(conn net.Conn, userName string) error {
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)
WriteAgain:
	now := time.Now()
	message := fmt.Sprintf("[%v][%v]:", now.Format("2006-01-02 15:04:05"), userName)
	_, err := conn.Write([]byte(message))
	if err != nil {
		return err
	}
	n, err2 := conn.Read(tmp)
	if err2 != nil {
		return err2
	}
	packet = packet[:0]
	packet = append(packet, tmp[:n]...)
	if !ValidInput(packet) {
		goto WriteAgain
	}
	fmt.Println(strings.Trim(string(packet), "\n"))
	goto WriteAgain
}
