**1)** Explain what is GO?

**Ans)** GO is an open source programming language which makes it easy to build simple, reliable and efficient software. Programs are constructed from packages, whose properties allow efficient management of dependencies.

**2)** What is syntax like in GO?

**Ans)** Syntax in GO is specified using Extended Backus-Naur Form (EBNF)

        Production = production_name “=” [ Expression ]
        Expression = Alternative { “l” Alternative }
        Alternative = Term { Term }
        Term = Production_name l token [ “…”token] l Group l Option l Repetition
        Group = “ ( “ Expression”)”
        Option = “ [ “ Expression “ ]”
        Repetition = “ {“ Expression “}”

**3)** Explain what is string literals?

**Ans)** A string literals represents a string constant obtained from concatenating a sequence of characters.

    There are two forms,
        A. Raw string literals: The value of raw string literals are character sequence between back quotes ‘‘.  The value of a string literal is the string composed of the uninterrupted character between quotes.
        B. Interpreted string literals: It is represented between double quotes ““. The text between the double quotes which may not contain newlines, forms the value of the literal.

**4)** Explain packages in Go program?

**Ans)** Every GO program is made up of packages. The program starts running in package main. This program is using the packages with import paths “fmt” and “math/rand”.

**5)** Explain workspace in GO?

**Ans)** Inside a workspace GO code must be kept. A workspace is a directory hierarchy with three directories at its root.

        src contains GO source files organized into packages
            pkg contains package objects and
            bin contains executable commands

**6)** What must be done to compile a program on a windows machine so that it
will run on Ubuntu Linux?
**Ans)** Set GOOS to "linux"

**7)** Which Go Keyword transfers the control of one switch case to the next case?
**Ans)** fallthrough

**8)** You are writing a small program and would like to log a file to detect races inside your current working directory $HOME/go/src/myprogram, rather than stderr. What would be the best way to make that work?

**Ans)** Set the GORACE environment variable to $HOME/go/src/myprogram

**9)** What is the correct way of removing an element from a slice?

**Ans)** a = append(a[:i], a[i+1]...)

**10)** In the `godoc` command, what should the value of the
`-index_throttle` flag be so index creation runs at full throttle?

**Ans)** 1.0

**11)** How can a set of subtests be configured to run in parallel? Assume that the "t" object is of type `*testsing.T`.

**Ans)** Add t.Parallel() call to each subtest.

**12)** What package contains functions that allow files to be opened on the filesystem?

**Ans)** `os`

**13)** What is panic in Golang?

**Ans)** A built-in function that stops teh function where it is called and any
deferred functions in the function are executed. Then, it returns back to the caller and continues until all functions are returned. At that point, the program terminates and prints the error message and stack trace.

**14)** Given below is a hash function. Which of the following pairs will have the same hash value? choose all that apply.

	hashFunction(string s){
		int hash = 0;
		for (int i=0;i<s.length;i++){
			hash += (i+1) * (s[i] - 'a' + 1);
		}
		return hash;
	}

	Example:
 
		s = 'abeaa'
		hash = 1*1 + 2*2 + 3*5 + 4*1 + 5 *1 = 29

	Pick ONE or MORE options
		a) "xwxx" and "vztz"
		b) "uwvy" and "gvzz"
		c) "tttt" and "zszt"
		d) "bvvv" and "xxxw"
---

    In mac, below command will create mac specific binary
    	$ GOOS=darwin GOARCH=386 go build someFile.go

    How to create such cross-platform builds in windows, command line

===

anonymous functions - it has access to data both inside
and outside of its code block - both local scope & the same scope
of outer function in which the closure is defined

package import format, if package is purely
needed for init function execution & not used in
in code. Used to import “side effects” (static reference) of a module

    import \_ "image/png"

`if` statements with initializers can be used
to create more than one variable and can
vary in type.

go build github.com/ps/foo/...
To compile the package "github.com/ps/foo" and all its sub-packages

go install ./...
from within parent package directory to compile a package and all od
its subpackages , placing the result in workspace's pkg directory

go install
will, by default, compile a library and add it to the
workspace's bin directory

Why Golang
--> For its Simplicity, Build in Concurrency Support and Super Scalable.
What is Goroutine
--> Goroutine is like a thread in java, it's lightweight than a thread useful to achieve concurrency in Go. https://golangbot.com/goroutines/
How to communicate between goroutines?
--> By Using Channels we can communicate between Goroutines.
What is mutex?
-- > https://golangbot.com/mutex/
What is synchronization in go
--> https://golang.org/pkg/sync/
How to achieve concurrency in Go
--> By Using Goroutines.
What is Channels in go?
--> https://golangbot.com/channels/
What are the Types of Channels?
--> https://golangbyexample.com/channel-golang/
What is Buffered Channel and Unbuffered Channel explain?
--> https://golangbot.com/channels/
How to avoid deadlock in a goroutine?
--> https://www.geeksforgeeks.org/golang-deadlock-and-default-case-in-select-statement/
How Goroutines Works?
--> https://golangbyexample.com/goroutines-golang/
What is a method in Go?
--> https://medium.com/rungo/anatomy-of-methods-in-go-f552aaa8ac4a
What is the anonymous function?
--> https://www.geeksforgeeks.org/anonymous-function-in-go-language/
What is the Variadic function?
--> https://www.geeksforgeeks.org/variadic-functions-in-go/
Explain Slices in Go in detail how they differ from another language?
--> https://golangbyexample.com/slice-in-golang/
What is Closures in go?
--> https://www.geeksforgeeks.org/closures-in-golang/
What is an interface in Go?
--> https://golangbyexample.com/interface-in-golang/
Go is Object-Oriented Langauge?
--> No
Explain the OOPs concept.
--> https://www.youtube.com/watch?v=pTB0EiLXUC8
https://golangbyexample.com/tag/oop/
How to create a basic server in Go?
--> https://gobyexample.com/http-servers
What are the basic data structures in Go?
--> https://medium.com/@victorsteven/understanding-data-structures

What is reflection
--> https://golangbot.com/reflection/
What is byte [] slice
--> https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go
How to create a Public/Private function in go (Access modifier in Go?)
--> We use Capital name(First Char Capital) for the function to make it Public and small letter for Private function.
Does go function return multiple values?
--> Yes, https://www.geeksforgeeks.org/golang-program-that-uses-multiple-return-values/
Difference between concurrency and Parralism?
--> https://golangbot.com/concurrency/
What is a map in go?
--> https://medium.com/rungo/the-anatomy-of-maps-in-go-79b82836838b
Difference between new and make?
--> https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives
What is Worker Pool and WaitGroup?
--> https://golangbot.com/buffered-channels-worker-pools/
How map prints data
--> https://golangbyexample.com/different-ways-iterating-over-map-go/
What is a rune?
--> https://www.bogotobogo.com/GoLang/GoLang_byte_and_rune.php
How Append function Works?
--> https://yourbasic.org/golang/append-explained/
What is the REST API?
--> https://www.smashingmagazine.com/2018/01/understanding-using-rest-api/
https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
What is Middleware?
--> https://www.redhat.com/en/topics/middleware/what-is-middleware
What is microservices?
--> https://www.redhat.com/en/topics/microservices/what-are-microservices
How to handle error in go?
--> https://medium.com/rungo/error-handling-in-go-f0125de052f0
What is the Defer keyword?
--> https://golangbot.com/defer/
What is Recover?
--> https://golangbot.com/panic-and-recover/
Difference between function and method?
--> https://www.sohamkamani.com/golang/functions-vs-methods/
What is the empty interface?
--> https://dev.to/flrnd/understanding-the-empty-interface-in-go-4652
What is a struct type?
--> https://golangbyexample.com/struct-in-golang-complete-guide/
What is Enum in go?
--> https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
What is Go Module?
--> go-modules
How to manage dependency in Go?
--> https://golangbyexample.com/packages-modules-go-first/
What is Packages in go?
--> https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go
What is Select in go?
--> https://golangbyexample.com/select-statement-golang/
What is IOTA?
--> https://golangbyexample.com/iota-in-golang/
What is panic in go how to handle it?
--> https://www.digitalocean.com/community/tutorials/handling-panics-in-go
Do you know net/http package?
--> https://codegangsta.gitbooks.io/building-web-apps-with-go/content/http_basics/index.html
What is JSON Encoding and Decoding?
--> https://yourbasic.org/golang/json-example/

Write any method code
Swap if the number is greater than the next number
Find the key from the array (not in any easy way, split the array into two then search)
Find the count of words in a long string
Check if the website is active or not?
Find the random number with x length
Find longest increasing subsequence in an array
Get unique values from a slice of int with count
Write palindrome program in go
Get pairs with x difference only unique pairs
Reverse a string Input = ‘Hello’, Output = ‘olleH’
Write Code :- Input : "aabbbcccab" , Output: "aa2bbb3ccc3a1b1"
Print unique elements form array, Input = [2,3,4,5,2,3], Output =  [2,3,4,5]
Find the lengthiest word.  Input =’ i am in Bangalore‘
Check whether it is closed or not.  Input =’{([])}’
Count the repeat numbers Input = [2,3,4,5,2,3,2, 5, 2], Output = {2:4, 3:2, 4:1, 5:2}
Display the middle number, Input = [2,3,4,5,2,3,7]
Find Combination of pair elements which give 0, Input = [ -1, 0, 1, -2, 2, 3, 4, 5, 2, 3]
Write a program use type interface with goroutines and channels.
Generate N digit random number based on input. input : 5 , Output : 93657 Input : 2, Output : 47

What is Package Context in Go" and how we are using it? 

What is the type Assertion?

How Go's interface is different from Java Interface

Design database design for particular system consider normalization

Inline function in Go?

Inline struct in Go?

What is GORM?

What is Gorilla Mux?

Difference between Switch and Select in Go?

What is First-Class Function in Go?

How to define an anonymous function in Go?

How to define an anonymous struct in Go?

Function closures in Go

What is Pass by Value and Pass by Reference?

What is Slice Capacity and Length?

What is struct composition?

What is a nested struct?

How to delete elements from Slice?

What is Nil Interface

What is Zero Value, Write down for all basic types?

What is Atoi and Itoa ?

What is gRPC?


Q) What will be the output of this program?

	ch := make(chan int)
	close(ch)
	val := <-ch 
	fmt.Println(val)
	
	a) NaN	b) It will panic 	c) It will deadlock 	d) 0
	
Q) What is the common way to have several executables in your project?
	a) Have a cmd directory and a directory per executable inside it 
	b) Have apkg directory and a directory per executable inside it 
	c) Use build tags 
	d) Comment out main 
	
Q) What is the difference between the time package's Time.Sub() and Time.Add() methods?
	a) Time.Add90 is for performing addition while Time.Sub() is for nesting timestamps. 
	b) They are opposites. Time.Add(x) is the equivalent of Time.Sub(-x)
	c) Time.Add() accepts a Duration parameter and returns a Time while Time.Sub() accepts a time parameter and returns a Duration. 
	d) Time.Add() always returns a later time while time.Sub always returns an earlier time. 
	
Q) What does a sync.Mutex block while it is locked?
	a) any other call to lock that Mutex 
	b) any writes to the variable it is looking 
	c) all goroutines 
	d) any reads or writes of the variable it is locking
