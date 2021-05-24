Go decided that the best place to handle an error is right there when you get it.

# Defer

    - Defer statement pushes a function call into a list.
    - The list of saved calls are executed after the surrounding function returns.
    - Executed in LIFO order
    - Generally, used for cleanup actions
    - Rules
        - A deferred function's arguments are evaluated when the defer statement is evaludated
        - Deferred function calls are executed in Last In First Out order after the surrounding function returns.
        - Deferred functions may read and assign to the returning function's named return values.

# Panic

    - built-in function that stops the ordinary flow of control and begins panicking.
    - They can be initiated by invoking panic directly.
    - They can also be caused by runtime errors, such as out-of-bounds array accesses.

# Recover

    - built-in function that regains control of a panicking goroutine.
    - Recover is only useful inside deferred functions.
    - During normal execution, a call to recover will return nil
      and have no other effect.
    - If the current goroutine is panicking, a call to recover
      will capture the value given to panic and resume normal execution.
