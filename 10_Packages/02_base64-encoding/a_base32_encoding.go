package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	originalMessage := "abc123!?$*&()'-=@~"
	fmt.Println("originalMessage=", originalMessage)

	encodedMessage := base32.StdEncoding.EncodeToString([]byte(originalMessage))
	fmt.Println("encodedMessage =", encodedMessage)

	decodedMessage, _ := base32.StdEncoding.DecodeString(encodedMessage)
	fmt.Println("decodedMessage =", string(decodedMessage))
}
