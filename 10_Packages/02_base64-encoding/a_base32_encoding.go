package main

import (
	"encoding/base32"
	"fmt"
	"reflect"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	orgMsg := "abc123!?$*&()'-=@~"
	fmt.Println("orgMsg=", orgMsg, reflect.TypeOf(orgMsg))

	byteStr := []byte(orgMsg)
	fmt.Println("byteStr=", byteStr, reflect.TypeOf(byteStr))

	encodeMsg := base32.StdEncoding.EncodeToString(byteStr)
	fmt.Println("encodeMsg=", encodeMsg, reflect.TypeOf(encodeMsg))

	decodeMsg, err := base32.StdEncoding.DecodeString(encodeMsg)
	CheckError(err)
	fmt.Println("decodeMsg=", decodeMsg, reflect.TypeOf(decodeMsg))
}
/*
The types of encoding available are:
	StdEncoding		: standard base64 encoding
	URLEncoding		: alternate base64 encoding used in URLs and file names
	RawStdEncoding	: standard, raw, unpadded base64 encoding
	RawURLEncoding	: unpadded alternate base64 encoding for URLs and file names
*/