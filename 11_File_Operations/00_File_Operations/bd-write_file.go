package main

import "io/ioutil"

func main() {
	message := []byte("first line of message")
	ioutil.WriteFile("anotherFile.txt", message, 0644)

}
