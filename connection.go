package main

import (
	"fmt"
	"strings"
)

const (
	tcpIsResolved int = 1 << iota
	tcpIsConnected
	tcpIsClosed
	udpIsResolved
	udpIsConnected
	udpIsClosed
)

func tryConnecting(host string, port string) int {
	addr := host + ":" + port
	return tcpToAddr(addr)
}

func getConnectionStatus(port string, result int) (string, bool) {
	status := "Port " + port + " is "
	protocol := ""
	isOpen := false

	if result&tcpIsConnected != 0 {
		isOpen = true
		protocol += ", tcp"
	}

	if result&udpIsConnected != 0 {
		isOpen = true
		protocol += ", udp"
	}

	if !isOpen {
		status += "closed"
	} else {
		tail, _ := strings.CutPrefix(protocol, ", ")
		status += "open"
		status += " (" + tail + ")"
	}

	return status, isOpen
}

func verifyConnection(host string, port string, showClosed bool) {
	result := tryConnecting(host, port)
	status, isOpen := getConnectionStatus(port, result)

	if !isOpen && !showClosed {
		return
	}
	fmt.Println(status)
}

// func verifyConnections(host string, ports []string, showClosed bool) {
// 	// for _, port := range ports {
// 	// 	go func(host string, port string, showClosed bool) {
// 	// 		verifyConnection(host, port, showClosed)
// 	// 	}(host, port, showClosed)
// 	// }
// }
