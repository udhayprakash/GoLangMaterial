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

/*
The types of encoding available are:
	StdEncoding		: standard base64 encoding
	URLEncoding		: alternate base64 encoding used in URLs and file names
	RawStdEncoding	: standard, raw, unpadded base64 encoding
	RawURLEncoding	: unpadded alternate base64 encoding for URLs and file names
*/
