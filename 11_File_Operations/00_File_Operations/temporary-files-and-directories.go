package main

/*

Throughout program execution, we often want to create data
that isn’t needed after the program exits. Temporary files
and directories are useful for this purpose since they don’t
pollute the file system over time.
*/
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := ioutil.TempFile("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	// Remember to clean up the file afterwards
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
