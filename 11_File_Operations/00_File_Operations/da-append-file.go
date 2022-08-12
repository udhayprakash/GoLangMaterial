package main

import "os"

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func appendFileData(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()

	_, err = f.WriteString(data)
}
func main() {

	// Trimming the data, in file, from start till the given position
	os.Truncate("testfile.txt", 10)
	appendFileData("testfile.txt", "Hello World!!!\n")

}
