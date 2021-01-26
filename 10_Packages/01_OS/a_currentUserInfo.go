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
	fmt.Println("Hi " + userObject.Name + " (id: " + userObject.Uid + ")")
	fmt.Println("Username: " + userObject.Username)
	fmt.Println("Name    : " + userObject.Name)
	fmt.Println("Home Dir: " + userObject.HomeDir)

	// Get "Real" User under sudo. - Useful in non-windows machines
	// More Info: https://stackoverflow.com/q/29733575/402585
	fmt.Println("Real User: " + os.Getenv("SUDO_USER"))

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("current working directory:", cwd)
}
