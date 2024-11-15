package functions

import (
	"io"
	"net"
	"strings"
)

func handleName(conn net.Conn, userName *string) error {
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)
	inputName := "[ENTER YOUR NAME]:"

check:
	_, err2 := conn.Write([]byte(inputName))
	if err2 != nil {
		return err2
	}
	n, err := conn.Read(tmp)
	if err != nil {
		if err == io.EOF {
			return err
		}
	}
	packet = packet[:0]
	packet = append(packet, tmp[:n]...)
	if (*userName) == "" {
		if !ValidInput(packet) {
			goto check
		}
		(*userName) = strings.Trim(string(packet), "\n")
	}

	return nil
}
