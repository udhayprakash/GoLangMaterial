package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {
	for _, ipAddr := range []string{
		"192.0.2.2",
		"292.0.2.2",
		"192.0.2.2456",
		"192.0.2.245",
	} {
		ip := net.ParseIP(ipAddr)
		if ip != nil {
			fmt.Println("Ip address: ", ip)
			fmt.Println("Type      : ", reflect.TypeOf(ip))
			fmt.Println("DefaultMask:", ip.DefaultMask())
			fmt.Println("isLoopBack: ", ip.IsLoopback())

		} else {
			fmt.Println("Invalid IP address:", ipAddr)
		}
		fmt.Println()
	}

}
