package main

import (
	"net"
)

func SendUDP(host, data string) {
	conn, _ := net.Dial("udp", host)
	if conn == nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(data))
}
