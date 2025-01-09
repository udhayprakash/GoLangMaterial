package main

import "fmt"

// go install go.uber.org/mock/mockgen@latest
// Mocking is done only for the interfaces, in Go

// func hello() {
// 	fmt.Println("Hello world")
// }


// Define an interface
type HelloService interface {
	Hello()
}

// Real implementation of the HelloService interface
type RealHelloService struct{}

func (s RealHelloService) Hello() {
	fmt.Println("Hello, World!")
}

// Standalone function that uses the HelloService interface
func hello(service HelloService) {
	service.Hello()
}

// mockgen -source=a_hello_world.go -destination=mock_hello.go -package=main
