## race condition

    - condition in which multiple goroutines are trying to access and modify the same resource.
    - It could be the case that one goroutines is trying to increase a value of a particular variable and the other goroutines is trying to access it at the same time, or it could be like multiple goroutines are trying to increase the value of a particular variable at the same time.

    - It should be noted that race conditions arise only if we have provided the writing permissions for a particular variable.
    - If only read permissions were available, then there wouldn't have been any problem as reading doesn't cause any issues, even if multiple goroutines are trying to read a single value.
