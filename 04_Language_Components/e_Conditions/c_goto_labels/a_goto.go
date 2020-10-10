package main

import "fmt"

func main() {
	fmt.Println("One")
	goto End
	fmt.Println("two")
	goto Udhay
End:
	fmt.Println("End label Block")
Udhay:
	fmt.Print("Udhay label block")
}
