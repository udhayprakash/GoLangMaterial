package main

import(
	"log"
	"os"
)

func main() {
	log.Println("message 1")

	file, e := os.OpenFile("04_LogRedirects.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}

	log.SetOutput(file)

	log.Println("message 2")

	// redirect logger to standard out
	log.SetOutput(os.Stdout)

	log.Println("message 3")

}