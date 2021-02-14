package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Here's the `string` we'll encode/decode.
	originalMessage := "abc123!?$*&()'-=@~"
	fmt.Println("originalMessage=", originalMessage)

	// Encode
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(originalMessage))
	fmt.Println("encodedMessage =", encodedMessage)

	// Decode
	decodedMessage, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		panic("malformed input")
	}
	fmt.Println("decodedMessage =", string(decodedMessage))
}
