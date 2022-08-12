package main

import (
	"fmt"
	"net"
)

func is_ip_addresses_equal(ip1, ip2 net.IP) bool {
	is_equal := ip1.Equal(ip2)
	return is_equal

}

func main() {
	ip1 := net.ParseIP("192.0.2.2")
	ip2 := net.ParseIP("192.0.2.02")

	if ip1 != nil && ip2 != nil {
		fmt.Println("Are both IPs equal: ", is_ip_addresses_equal(ip1, ip2))
	} else {
		fmt.Println("Invalid IP address. IP1: ", ip1, "IP2: ", ip2)
	}
}
