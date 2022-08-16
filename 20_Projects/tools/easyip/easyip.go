package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	quieter := flag.Bool("q", false, "Quieter, print only the IP. (Except for errors)")
	newline := flag.Bool("n", false, "Don't print a line feed at the end. (Only affects Unix systems)")
	info := flag.Bool("i", false, "Print information about the script.")
	flag.Parse()

	if *info {
		fmt.Println("\nHosts checked (in order, in case the previous one fails):")
		fmt.Println(" - http://ifconfig.me/ip\n - http://api.ipify.org\n - http://ipecho.net/plain\n - https://ip.seeip.org")

		return
	}

	if !*quieter {
		fmt.Print("Public IP: ")
	}
	res, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		res, err = http.Get("http://api.ipify.org")
		if err != nil {
			res, err = http.Get("http://ipecho.net/plain")
			if err != nil {
				res, err = http.Get("https://ip.seeip.org")
				if err != nil {
					fmt.Print("\rUnable to fetch public IP.")
					os.Exit(1)
				}
			}
		}
	}
	ip, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Print("\rUnable to fetch public IP.")
		os.Exit(1)
	}
	fmt.Printf("%s", ip)
	if !*newline {
		fmt.Println()
	}
}
