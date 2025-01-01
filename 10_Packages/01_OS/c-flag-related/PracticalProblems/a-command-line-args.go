package main

import (
	"fmt"
	"os"
	"strings"
)

var username, password, servername string

func init() {

	// // Method 1 - hard coding
	// username = "ravi"
	// password = "pass@123"
	// servername = "something.company.com"

	// // Method 2 - getting details in runtime
	// fmt.Print("Enter username:")
	// fmt.Scanln(&username)
	// fmt.Print("Enter password:")
	// fmt.Scanln(&password)
	// fmt.Print("Enter server Name:")
	// fmt.Scanln(&servername)

	// Method 3 - command line arguments
	// 	fmt.Println("os.Args:", os.Args)

	// 	fmt.Println("No of args:", len(os.Args))

	const usage = `USAGE: %s <username> <password> <servername>
	Example:
	%s ravi pass@123 something@company.com
	`
	if len(os.Args) != 4 {
		// fmt.Println("USAGE: filename.go <username> <password> <servername>")
		fmt.Printf(usage, os.Args[0], os.Args[0])
		os.Exit(1)
	}

	// 	fmt.Println("this file :", os.Args[0])
	// 	fmt.Println("Arguments:", os.Args[1:])
	// 	fmt.Println(reflect.TypeOf(os.Args)) // []string

	username = os.Args[1]
	password = os.Args[2]
	servername = os.Args[3]
}

func main() {

	// Validate arguments
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" || strings.TrimSpace(servername) == "" {
		fmt.Println("Error: username, password, and servername cannot be empty")
		os.Exit(1)
	}

	// Print the values (mask the password for security)
	maskedPassword := strings.Repeat("*", len(password))
	fmt.Printf(`
		username: %s
		password: %s
		servername: %s
		
	`, username, maskedPassword, servername)

}
