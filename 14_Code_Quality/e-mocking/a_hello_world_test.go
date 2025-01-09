// a_hello_world_test.go
package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of the HelloService interface
	mockHelloService := NewMockHelloService(ctrl)

	// Set expectations on the mock
	mockHelloService.EXPECT().Hello().Do(func() {
		fmt.Println("Mock Hello, World!")
	})

	// Call the hello function with the mock
	hello(mockHelloService)
}
