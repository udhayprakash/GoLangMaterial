package main

/*
 #include <stdio.h>

void myprint(char* s) {
	printf("%s\n", s)
}

*/
import "C"

func main() {
	C.myprint(C.CString("Hello world"))
}
