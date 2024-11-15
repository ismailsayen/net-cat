package utils

import "net"

func Writer(s string, conn net.Conn) error {
	if _, err := conn.Write([]byte(s)); err != nil {
		return err
	}

	return nil
}
