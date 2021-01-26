package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Here's the `string` we'll encode/decode.
	originalMessage := "abc123!?$*&()'-=@~"
	fmt.Println("originalMessage=", originalMessage)

	encodedMessage := base64.StdEncoding.EncodeToString([]byte(originalMessage))
	fmt.Println("encodedMessage =", encodedMessage)

	decodedMessage, _ := base64.StdEncoding.DecodeString(encodedMessage)
	fmt.Println("decodedMessage =", string(decodedMessage))

	// ENCODING for URLs
	encodedMessage = base64.URLEncoding.EncodeToString([]byte(originalMessage))
	fmt.Println("\nencodedMessage =", encodedMessage)

	decodedMessage, _ = base64.URLEncoding.DecodeString(encodedMessage)
	fmt.Println("decodedMessage =", string(decodedMessage))
}

//YWJjMTIzIT8kKiYoKSctPUB+
//YWJjMTIzIT8kKiYoKSctPUB%2B
