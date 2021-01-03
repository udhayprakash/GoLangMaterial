package main

 import (
 	"crypto/md5"
 	"fmt"
 	"io"
 	"os"
 )

 func main() {
 	file, err := os.Open("utf8.txt")

 	if err != nil {
 		panic(err)
 	}

 	defer file.Close()

 	hash := md5.New()
 	_, err = io.Copy(hash, file)

 	if err != nil {
 		panic(err)
 	}

 	fmt.Printf("%s MD5 checksum is %x \n", file.Name(), hash.Sum(nil))

 }