package main

import (
	"fmt"
	"log"
	"os"
)

func CanWrite(filepath string) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false, err
		}
	}
	file.Close()
	return true, nil

}

func SetWritable(filepath string) error {
	err := os.Chmod(filepath, 0222)
	return err
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	return err
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage : %s <filename>\n", os.Args[0])
		os.Exit(0)
	}

	filename := os.Args[1]

	fmt.Printf("Set %s to read only mode\n", filename)
	err := SetReadOnly(filename)

	if err != nil {
		log.Fatal(err)
	}

	// test after set to READ ONLY
	result, err := CanWrite(filename)
	if err != nil {
		//fmt.Println(err) -- permission denied because of readonly
	}
	fmt.Printf("Is %s writable? : [%v]\n", filename, result)

	fmt.Println("Set", filename, " to writable")
	err = SetWritable(filename)
	if err != nil {
		log.Fatal(err)
	}

	// test after set to WRITE ONLY
	result, err = CanWrite(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Is %s writable? : [%v]\n", filename, result)
}
