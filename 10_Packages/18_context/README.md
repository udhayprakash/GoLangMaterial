# go-context

Across the goroutines, it Manages
    
    - request-scoped values
    - cancellation signals
    - timeouts

Values can be passed to functions 
    (traditional functions)
        1)  by arguments (call by value), or with pointer args (call by reference)
    (go routines)
        2) by channels (only to one specific function, single or bidirectional)
        3) using context (multiple functions, read or write)


## Basic context functionality
`context.Background()` is root context in Go, which is 

    - empty context that is never cancelled, has no deadline, and carries no value

`context.TODO()` is placeholder context, used when we plan to replace in later stage of code, with proper context
    - same like context.Background
    - But signals to developers that context is not yet finalized

`context.WithCancel(parent Context)`
    Creates a new context that can be cancelled using the returned cancel function.

`context.WithTimeout(parent Context, timeout time.Duration)`:
    Creates a new context that automatically cancels after the specified timeout.

`context.WithDeadline(parent Context, deadline time.Time)`:
    Creates a new context that automatically cancels at the specified deadline.


`context.WithValue(parent Context, key, val interface{}):`
    Creates a new context that carries a key-value pair.


## Types of contexts

### Request Cancellation:

Propagate cancellation signals across goroutines and functions.

`Example:` Cancelling a long-running operation when a user cancels a request.

`Metaphor`: A teacher (main goroutine) tells students (child goroutines) to stop working on their assignments because the class is over.
    
    Without context: Students keep working even after the class ends, wasting time and resources.

    With context: Students stop immediately when the teacher signals them.

### Timeout Management:

Set deadlines for operations to ensure they don’t run indefinitely.

`Example:` Timing out a database query if it takes too long.

`Metaphor:` A chef (main goroutine) sets a timer for baking a cake (operation). If the timer goes off, the chef stops baking, even if the cake isn’t fully done.

    Without context: The chef keeps baking indefinitely, potentially burning the cake.

    With context: The chef stops baking when the timer goes off.

### Deadline Propagation:

Pass deadlines through function calls and goroutines.

`Example:` Ensuring all parts of a request respect a global deadline.

`Metaphor:` A project manager (main goroutine) sets a deadline for a project and ensures all team members (child goroutines) respect it.

    Without context: Team members work without knowing the deadline, leading to delays.

    With context: Team members stop working when the deadline is reached.

### Value Propagation:

Pass request-scoped values (e.g., user IDs, request IDs) across functions and goroutines.

`Example:` Attaching a trace ID for distributed tracing.

`Metaphor:` A librarian (main goroutine) gives a book (request-scoped value) to a student (child goroutine) and ensures they return it with their notes.

    Without context: The student loses track of the book and can’t return it.

    With context: The student returns the book with their notes, and the librarian knows who used it

### Graceful Shutdown:

Signal goroutines to stop work during application shutdown.

`Example:` Stopping background workers when the server is shutting down.

`Metaphor:` A conductor (main goroutine) signals musicians (child goroutines) to stop playing when the concert ends.

    Without context: Musicians keep playing even after the concert is over.

    With context: Musicians stop playing when the conductor signals them.

### Concurrency Control:

Manage goroutines and ensure they respect cancellation or timeouts.

`Example:` Running multiple tasks in parallel but stopping them if one fails.

`Metaphor:` A traffic light (main goroutine) controls the flow of cars (goroutines) at an intersection.

    Without context: Cars keep moving without coordination, leading to accidents.

    With context: Cars stop when the traffic light turns red

