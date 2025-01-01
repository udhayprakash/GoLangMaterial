package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("anotherfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	_, err = f.Write([]byte("Hello\n"))
	checkErr(err)

	f.Close()
	fmt.Println("FIle created and appended content")
}

/*
------------------------------------------------------------------
Flag		Meaning
------------------------------------------------------------------
O_RDONLY	It opens the file read-only
O_WRONLY	It opens the file write-only
O_RDWR		It opens the file read-write
O_APPEND	It appends data to the file when writing
O_CREATE	It creates a new file if none exists
O_EXCL		It is used with O_CREATE, and the file must not exist
O_SYNC		It opens for synchronous I/O
O_TRUNC		If possible, truncate file when opened
-------------------------------------------------------------------


*/
