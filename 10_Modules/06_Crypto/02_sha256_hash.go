package main

import (
	"crypto/sha256"
	"fmt"
)

func bitDifference(digest1, digest2 [32]byte) (count int) {
	fmt.Println("len(digest1):", len(digest1))
	fmt.Println("len(digest2):", len(digest2))
	for index1, eachChar1 := range digest1 {
		if eachChar1 != digest2[index1] {
			count++
		}
	}
	return count
}
func main() {
	// SHA256 cryptographic hash or digest of message is stored
	// in an arbitrary byte slice
	c1 := sha256.Sum256([]byte("x")) // lower-case
	c2 := sha256.Sum256([]byte("X")) // upper-case
	fmt.Printf("%x\n%x\n%t\n", c1, c2, c1 == c2)
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	fmt.Printf("%T\n", c1) // [32]uint8
	fmt.Println("bitDifference(c1, c2):", bitDifference(c1, c2))

	// NOTE: The two inputs differ by only a single bit, but
	//   approximately half the bits are different in the digests.
}
