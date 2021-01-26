package main

func main() {
	// calling a function
	helloWorld()

	// creating a go-routine
	go helloWorld()
}

func helloWorld(){
	println("Hello World")
}