package main

import (
	"net"
)

func tcpToAddr(addr string) int {
	result := 0

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return result
	}

	result |= tcpIsResolved

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		result |= tcpIsClosed
		return result
	}

	result |= tcpIsConnected

	conn.Close()

	return result
}
