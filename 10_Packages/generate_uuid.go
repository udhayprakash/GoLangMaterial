package main

import (
	"encoding/ascii85"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	originalUUID := uuid.New()
	originalSlice, _ := originalUUID.MarshalBinary()
	enc := make([]byte, ascii85.MaxEncodedLen(len(originalUUID)))
	ascii85.Encode(enc, originalSlice)
	newStringID := originalUUID.String()

	fmt.Printf("uuid    - len: %d, string: '%s'\n", len(newStringID), newStringID)
	fmt.Printf("ascii85 - len: %d, string: '%s'\n", len(enc), string(enc))

	// Just for testing:
	dst := make([]byte, 16)
	ascii85.Decode(dst, enc, true)
	uuidDecoded, _ := uuid.FromBytes(dst)
	fmt.Printf("decoded - len: %d, string: '%s'\n", len(uuidDecoded.String()), uuidDecoded.String())
}
