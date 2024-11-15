package models

import "net"

type Users struct {
	List map[string]net.Conn
}
