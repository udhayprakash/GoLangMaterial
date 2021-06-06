package main

import (
	"fmt"
	"net"
	"net/url"
)

func checkStringForIpOrHostname(host string) {
	addr := net.ParseIP(host)
	if addr == nil {
		fmt.Println("Given String is a Domain Name")

	} else {
		fmt.Println("Given String is a Ip Address:", addr)
	}
}
func main() {
	checkStringForIpOrHostname("google.com")
	checkStringForIpOrHostname("192.168.1.0")

	// Get Domain Name from address
	urlInstance, error := url.Parse("https://mydomain.com/category/mypage.html")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(urlInstance.Hostname())

	// Get Local IP address of a system
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)

	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(ipAddress)
}
