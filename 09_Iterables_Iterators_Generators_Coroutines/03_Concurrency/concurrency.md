Concurrency
-----------
“Concurrency is about dealing with lots of things at once. Parallelism
is about doing lots of things at once” – Rob Pike
- Two types in Go
    1) Communicating Sequential Processes (CSP)
        - Using communication as synchronization primitive
        - Goroutines
    2) Shared Memory Multi-threading 
        - Using locks
        
        
Go-Routines
-----------
- main() function runs in main routine.
- when main() retuns, all goroutines terminate.

Concurrency with Shared Memory
------------------------------
    sync.Mutex      Mutual exclusion with lock/unlock 
    sync.RWMutex    Multiple read, single write
    sync.Once       Initialize variables once
         
        

Useful Links
-------------
1) Concurrency via partial ordering(happens-before order) https://golang.org/ref/mem