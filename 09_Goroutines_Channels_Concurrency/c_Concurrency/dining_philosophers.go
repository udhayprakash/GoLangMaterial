package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	leftChopStick, rightChopStick *ChopStick
	position int
	hostChannel chan int
}

func main() {

	// Create a buffered channel, this allows a maximum of two
	// philosophers to eat at the same time
	hostChannel := make(chan int, 2)

	ChopSticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	Philosophers := make([]*Philosopher, 5)

	for j := 0; j < 5; j++ {
		Philosophers[j] = &Philosopher{ChopSticks[j], ChopSticks[(j+1)%5], j+1, hostChannel}
	}

	var waitGroup sync.WaitGroup

	for k := 0; k < 5; k++ {

		waitGroup.Add(1)
		go func(index int) {
			Philosophers[index].Eat()
			waitGroup.Done()
		}(k)

	}
	waitGroup.Wait()
}

func (p Philosopher) Eat() {

	for turns := 0; turns < 3; turns++ {

		// This is a buffered channel with capacity 2
		// If there are already two philosophers eating then this statement will block
		p.hostChannel <- 1

		fmt.Printf("Starting to eat %d (turn %d)...\n", p.position, turns+1)

		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		p.rightChopStick.Unlock()
		p.leftChopStick.Unlock()

		fmt.Printf("Finishing eating %d (turn %d)...\n", p.position, turns+1)

		// Release a spot for other philosophers to start eating
		<- p.hostChannel
	}
}