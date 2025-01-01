package main

import (
	"flag"
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

	/*

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
	*/

	// Method 4 - using flag package
	// Define flags
	// username := flag.String("username", "", "Username for authentication")
	// password := flag.String("password", "", "Password for authentication")
	// servername := flag.String("servername", "", "Server name to connect to")

	// To use package-level variables
	flag.StringVar(&username, "username", "", "Username for authentication")
	flag.StringVar(&password, "password", "", "Password for authentication")
	flag.StringVar(&servername, "servername", "", "Server name to connect to")

}

func main() {

	// Parse flags
	flag.Parse()

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

	connectToServer()

}

// Another function that uses the package-level flag variables
func connectToServer() {
	fmt.Printf("\nConnecting to server %s with username %s...\n", servername, username)
}

// $ go run 01-command-line-args-using-flags.go -h

// $ go run 01-command-line-args-using-flags.go -username ravi -password pass@123 -servername something.company.com
