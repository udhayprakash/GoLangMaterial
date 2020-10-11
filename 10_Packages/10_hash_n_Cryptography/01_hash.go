package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// constants
	fmt.Println("crc32.Size      =", crc32.Size)
	fmt.Println("crc32.Castagnoli=", crc32.Castagnoli)
	fmt.Println("crc32.Koopman   =", crc32.Koopman)
	fmt.Println("crc32.IEEE      =", crc32.IEEE)
	fmt.Println("crc32.IEEETable =", crc32.IEEETable)

	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("Thi is test data"))

	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println("v=", v)

}
