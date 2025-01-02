package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	var port string
	var showClosed bool
	var maxConnections int

	flag.StringVar(&host, "host", "0.0.0.0", "host to connect to")
	flag.StringVar(&port, "port", "1-65535", "port to connect to. Usage: -port 80 or -port 80,443,8080 or -port 1-65535")
	flag.BoolVar(&showClosed, "show-closed", false, "show closed ports")
	flag.IntVar(&maxConnections, "max-con", 0, "maximum number of connections to make concurrently")

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

	if maxConnections > 0 {
		blocker := make(chan struct{}, maxConnections)
		for _, port := range ports {
			blocker <- struct{}{}
			go func(host string, port string, showClosed bool) {
				verifyConnection(host, port, showClosed)
				<-blocker
			}(host, port, showClosed)
		}
		return
	}

	for _, port := range ports {
		go verifyConnection(host, port, showClosed)
	}
}
