package main

import (
	"net"
)

func SendTCP(host, data string) {
	conn, _ := net.Dial("tcp", host)
	if conn == nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(data))
}
