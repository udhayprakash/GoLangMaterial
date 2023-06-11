package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("καλημ ρα κóσμ 你好, नमस्ते, Привет, ᎣᏏᏲ") // Go supports UTF-8 by default.
}

/*
Go commands
-----------
	To run the program,
		go run a-hello-world.go

		-work option prints the working directory path in addition to running the program, but it doesn't delete the working directory afterward
		go run - work a-hello-world.go


	To display the steps, instead of actually creating binary,
		go build -n a-hello-world.go

	To display all steps involved in binary creation,
		go build -x a-hello-world.go

	To build the executable (with same name),
		go build a-hello-world.go

	To build the executable (with specific name),
		go build -o Greetings.exe a-hello-world.go

	To build Executable for different operating systems,
		set GOOS=linux
		set GOARCH=amd64
		go build -o hello_linux_amd64 a-hello-world.go

		set GOOS=darwin
		set GOARCH=amd64
		go build -o hello_macos_amd64 a-hello-world.go

		set GOOS=windows
		set GOARCH=amd64
		go build -o hello_windows_amd64.exe a-hello-world.go


*/
