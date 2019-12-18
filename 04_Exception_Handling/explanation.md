Panic
=====
    - built-in function that stops the ordinary flow of control and begins panicking.
    - They can be initiated by invoking panic directly. 
    - They can also be caused by runtime errors, such as out-of-bounds array accesses.
    
Recover
=======
    - built-in function that regains control of a panicking goroutine. 
    - Recover is only useful inside deferred functions. 
    - During normal execution, a call to recover will return nil 
      and have no other effect. 
    - If the current goroutine is panicking, a call to recover 
      will capture the value given to panic and resume normal execution.
 
