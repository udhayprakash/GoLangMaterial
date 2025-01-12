package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, iface := range ifaces {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				fmt.Println("Local IP:", ipNet.IP.String())
			}
		}
	}
}