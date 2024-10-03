package main

import (
	"net"
)

// UDP is not a reliable for open connect checking
func udpToAddr(addr string) int {
	result := 0

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return result
	}

	result |= tcpIsResolved

	conn, err := net.DialUDP("tcp", nil, udpAddr)
	if err != nil {
		result |= tcpIsClosed
		return result
	}

	result |= tcpIsConnected

	conn.Close()

	return result
}
