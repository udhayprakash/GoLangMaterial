package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	//fmt.Println("user.Current():", user.Current())
	//multiple-value user.Current() in single-value context

	userObject, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Printf("Hi %s (id: %s)\n", userObject.Name, userObject.Uid)
	fmt.Println("Username: " + userObject.Username)
	fmt.Println("Name    : " + userObject.Name)
	fmt.Println("Home Dir: " + userObject.HomeDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fmt.Println("Home Dir:", homeDir)

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("HostName:", hostName)

	// Get "Real" User under sudo. - Useful in non-windows machines
	// More Info: https://stackoverflow.com/q/29733575/402585
	fmt.Println("\nReal User: " + os.Getenv("SUDO_USER"))
	fmt.Println("Real User: " + os.Getenv("USERNAME"))

}
