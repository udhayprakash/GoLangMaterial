package main

import (
	"fmt"
	"net"
)

func main() {
	// Validating IP address
	// ip := net.ParseIP("127..0.0.1") // false
	// ip := net.ParseIP("127.0.0.1")  // true
	ip := net.ParseIP("192.9.0.1") // false
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}
	fmt.Println("IP:", ip)
	fmt.Println("Is loopback:", ip.IsLoopback())
	fmt.Println()

	//  DNS Lookup - Resolve domain names to IP addresses and vice versa
	ips, err := net.LookupIP("google.com")
	if err != nil {
		fmt.Println("DNS lookup failed:", err)
		return
	}
	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}

}
