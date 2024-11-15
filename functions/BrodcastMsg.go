package functions

import (
	"fmt"
	"net"
	"time"
)

func BrodcastMsg(msg string, conn net.Conn) error {
	for userName, val := range list.List {
		if val != conn {
			message := fmt.Sprintf("[%v][%v]:", time.Now().Format(time.DateTime), userName)

			if _, err := val.Write([]byte(msg + message)); err != nil {
				return err
			}
		}
	}
	return nil
}
