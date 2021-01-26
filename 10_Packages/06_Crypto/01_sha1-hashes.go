package main

/*

SHA1 hashes are frequently used to compute short
identities for binary or text blobs.
For example, the git revision control system uses
SHA1s extensively to identify versioned files and directories.
*/
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
