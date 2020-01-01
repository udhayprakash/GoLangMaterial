Channels
--------
    - If goroutines are the activities of a concurrent Go program, channels are the 
      connections between them.
    - channels let one goroutine send values to another
            ch := make(chan int)    // unbuffered channel
            ch := make(chan int, 0) // unbuffered channel
            ch := make(chan int, 3) // buffered channel with capacity 3
    -  send: 
            ch <- x  // send value x to channel, ch
    - receive: 
            x = <-ch // receive value from channel, ch, and assign to x
            <-ch     // receive value from channel, ch, and discard value
    - close: 
        close(ch)    // closes channel, ch
      - sets a flag, after which 
          – Additional receives get zero value
          – Additional sends panic
      
    - Types
        1) Unbuffered (or synchronous) Channel
            - Send operation blocks sending goroutine unitl another goroutine 
              executes corrsponding receive on the same channel.
            - If receive operation was attempted first, receiving goroutine is
              blocked unit another goroutine performs a send on same channel.
            - Unbuffered channels "synchronize" sending and receiving goroutines.
            - When a value is sent, receipt of value happens before reawakening
              of the sending goroutine.
            
        2) Buffered Channel
            - Unidirectional buffers specify buffers as just senders
                - Receive-only
                    ch := make(<-chan int)
                - Send-only
                    ch := make(chan<- int)
            - If there are items in buffer, neither sender nor receiver are blocked
            - When
                buffer empty -> receiver blocked
                buffer full  -> sender blocked
            - Buffer Size should be choosen appropriately
                - Else deadlock or waiting of pipelined processes takes place.
        3) New Channel
                
    Closing a channel is a signal to its receivers that it is no longer accepting new data; nothing more, nothing less.
    Closing a channel is not necessary to "free" a channel. You don’t need to close a channel to "clean up" its resources.
        
    Channel pipelining
        - channels can be used to connect goroutines together so that output of one is input to another
            