package main

import (
	"fmt"
	"net"
)

func getIPClass(ip net.IP) string {
	if ip.To4() != nil { // Check if it's an IPv4 address
		fmt.Println(ip, "--", ip[0])
		firstOctet := ip[0]
		switch {
		case firstOctet >= 1 && firstOctet <= 126:
			return "Class A"
		case firstOctet >= 128 && firstOctet <= 191:
			return "Class B"
		case firstOctet >= 192 && firstOctet <= 223:
			return "Class C"
		case firstOctet >= 224 && firstOctet <= 239:
			return "Class D (Multicast)"
		case firstOctet >= 240 && firstOctet <= 255:
			return "Class E (Reserved)"
		}
	}
	return "Unknown Class"
}

func main() {
	for _, ipAddr := range []string{
		"192.0.2.2",
		"292.0.2.2",
		"192.0.2.2456",
		"192.0.2.245",
		"142.250.195.238",
		"10.0.0.1",        // Class A
		"172.16.0.1",      // Class B
		"192.168.1.1",     // Class C
		"224.0.0.1",       // Class D
		"240.0.0.1",       // Class E
		"255.255.255.255", // Broadcast address
		"0.0.0.0",         // Default route
	} {
		// String to IP address parsing
		ip := net.ParseIP(ipAddr)
		fmt.Printf(" %15v ==> %v \n", ipAddr, getIPClass(ip))

	}
}
