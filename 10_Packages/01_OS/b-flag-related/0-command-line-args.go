package main

import (
	"fmt"
	"os"
	"reflect"
)

func init() {
	if len(os.Args) != 4 {
		fmt.Println("USAGE: filename.go <username> <password> <servername>")
		os.Exit(1)
	}

}

func main() {
	var username, password, servername string

	// // Method 1 - hard coding
	// username = "ravi"
	// password = "pass@123"
	// servername = "something@tcs.com"

	// // Method 2 - getting details in runtime
	// fmt.Print("Enter username:")
	// fmt.Scanln(&username)
	// fmt.Print("Enter password:")
	// fmt.Scanln(&password)
	// fmt.Print("Enter server Name:")
	// fmt.Scanln(&servername)

	// Method 3 - command line arguments
	fmt.Println("os.Args:", os.Args)

	fmt.Println("No of args:", len(os.Args))
	fmt.Println("this file :", os.Args[0])
	fmt.Println("Arguments:", os.Args[1:])
	fmt.Println(reflect.TypeOf(os.Args)) // []string

	username = os.Args[1]
	password = os.Args[2]
	servername = os.Args[3]

	fmt.Printf(`
	username: %s
	password: %s
	servername: %s
	`, username, password, servername)
}
