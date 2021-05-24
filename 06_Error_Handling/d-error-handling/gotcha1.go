package main

func main() {
	message := make(chan string)
	select {
	case <-messages:
	}
}
/*
NOTE: 
It is missing default case in the "select"
statement casuing the "select" statement 
to be vlocked forever becaise no other goroutines
are writing to this channel.
*/