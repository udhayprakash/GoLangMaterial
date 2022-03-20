package main

/*
A goroutine has a simple model: it is a function executing
concurrently with other goroutines in the same address space.
It is lightweight, costing little more than the allocation of
 stack space. And the stacks start small, so they are cheap,
  and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one
should block, such as while waiting for I/O, others continue
to run. Their design hides many of the complexities of thread
creation and management.
*/

func helloWorld() {
	println("Hello world")
}

func main() {

	// call a function
	helloWorld() // Hello world

	// creating a go-routine
	go helloWorld()

}
