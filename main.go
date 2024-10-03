package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
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

func parsePorts(input string) ([]string, error) {
	ports := []string{}

	if input == "" {
		return ports, nil
	}

	input = strings.ReplaceAll(input, " ", "")

	parts := strings.Split(input, ",")

	for _, part := range parts {
		if strings.ContainsAny(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return ports, errors.New("invalid range")
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return ports, errors.New("invalid range")
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return ports, errors.New("invalid range")
			}

			if start > end {
				return ports, errors.New("invalid range")
			}

			for i := start; i <= end; i++ {
				ports = append(ports, strconv.Itoa(i))
			}
		} else {
			ports = append(ports, part)
		}
	}

	return ports, nil
}

func main() {
	var host string
	var port string
	var showClosed bool

	flag.StringVar(&host, "host", "", "host to connect to")
	flag.StringVar(&port, "port", "", "port to connect to. Usage: -port 80 or -port 80,443,8080 or -port 1-65535")
	flag.BoolVar(&showClosed, "show-closed", false, "show closed ports")

	flag.Parse()

	if host == "" {
		flag.PrintDefaults()
		return
	}

	if port == "" {
		flag.PrintDefaults()
		return
	}

	ports, err := parsePorts(port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, port := range ports {
		verifyConnection(host, port, showClosed)
	}
}
