package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	out, _ := os.Create("novel.txt") //give any name
	defer out.Close()

	resp, _ := http.Get("http://www.gutenberg.org/files/18581/18581.txt")
	defer resp.Body.Close()

	// copies the bytes from your download into the empty contents of the file created
	contentLength, _ := io.Copy(out, resp.Body)
	fmt.Println(contentLength) // 325146
}
