## Concurrency

“Concurrency is about dealing with lots of things at once. Parallelism
is about doing lots of things at once” – Rob Pike

- With no shared data, concurrence gets a lot simpler:
  No semaphores.
  No monitors.
  No locks.
  No race-conditions.
  No dead-locks.

- A process is running a computer program that may have multiple threads.
- concurrent program
  - When a process runs multiple threads on a single core/processor
- parallelized program
  - When a process runs multiple threads on multiple cores/processors
- In Golang, we can use goroutines to write a concurrent program and channel for communication
  between goroutines.

## Go's Solution for Concurrency

1. GoRoutines (execution)
   A GoRoutine in the Go programming language is a lightweight
   thread that is managed by Go runtime. If you just put 'go'
   before a function, it means that it will execute concurrently
   with the rest of the code.
2. Channels (communication)
   Channels are pipes that connect concurrent GoRoutines. You
   are able to send values and signals over Channels from GoRoutine
   to GoRoutine. This allows for synchronizing execution.
3. Select (coordination)
   The Select statement in Go lets you wait and watch multiple
   operations on a channel. Combining GoRoutines and channels will
   show off the true power of concurrency in Go.

## Go-Routines

- main() function runs in main routine.
- when main() returns, all goroutines terminate.

- Two types in Go
  1. Communicating Sequential Processes (CSP)
     - Using communication as synchronization primitive
     - Goroutines
  2. Shared Memory Multi-threading
     - Using locks

## Concurrency with Shared Memory

    sync.Mutex      Mutual exclusion with lock/unlock
    sync.RWMutex    Multiple read, single write
    sync.Once       Initialize variables once

## Useful Links

1. Concurrency via partial ordering(happens-before order) https://golang.org/ref/mem
