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
