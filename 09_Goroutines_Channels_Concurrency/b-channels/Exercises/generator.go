/*
Generator returns a channel that produces
the numbers 1, 2, 3, ...

To Stop the underlying goroutine, send a number on this channel.
*/

func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println("<-number", <-number) // 1
	fmt.Println("<-number", <-number) // 2

	number <- 0                       // Stops underlying goroutine
	fmt.Println("<-number", <-number) // error, no one is sending anymore
	// fatal error: all goroutines are asleep - deadlock!
}