# Go Lang Material

[![Udhay's GitHub activity graph](https://img.shields.io/github/commit-activity/m/udhayprakash/GoLangMaterial?label=Commit%20Activity&logo=github&style=flat-square)](https://github.com/udhayprakash/GoLangMaterial)

- General Purpose Compiled Programing language
- created by Ken Thompson (B, C, Unix, UTF-8), Rob Pike(Unix, UTF-8),
    Robert Griesemer(Hotspot, JVM) and few google engineers.
- first release was on March 2012.
- No pre-requisites need. Can be learned as first programming language.
- C language is the most direct influence on Go.
- Platform-independent.
    - Create executable build based on the platform.
    - Means, source code is translated into language that the system understands.
- comes with its own compiler
- statically typed programming language
- Concurrency
    - allows multiple processes running simultaneously and effectively.
- garbage collector - automatic memory management
- Execution flow - top to botton and left to right order
- Official website: https://www.golang.org
- Go programs are two types
    - Executables
    - Libraries
        - Libraries are collections of code that we package together so that we can use them in other programs.
- Go doesn't care about white-spaces and indentation
- Two type of comments are supported:
    - line comment //
    - multi-line comment /* */
- Functions are building blocks of a Go Program
- All versions are backward compatible (promised).
- Go has replaced the traditional features of class inheritance with the help of
its two features, i.e., type embedding and interfaces.
- Other features
    - Fast compilation process
    - Default concurrency support
    - Simple and easy to use
- Designed to achieve concurrency using the multi-core CPUs available these days
- compiles, combining all dependency libraries and module into a single binary file based
    on OS type and architecture
- Popular features of other languages, which are not present in Golang are
        - no implicit numeric conversions
        - no constructors or destructors
        - no operator overloading
        - no default parameter values
        - no inheritance
        - no generics
        - no exceptions
        - no macros
        - no function annotations
        - no thread-local storage
-  It is meant to replace C++ and Java in terms of Google's needs.
    Go was meant to alleviate some of the slowness and clumsiness of development
    of very large software systems.
- To be a little more specific, Go helps solve …
        - slow compilation and slow execution
        - programmers that collaborate using different subsets of languages
        - readability and documentation
        - language consistency
        - versioning issues
        - multi-language builds
        - dependencies being hard to maintain
- Compilers
    - There are two Go compiler implementations, gc, and gccgo.
        Gc uses a different calling convention and linker and because of this,
        can only be linked with C programs following the same convention.
        Gccgo is a GCC frontend that can be linked with GCC-compiled C or C++ programs.
        Gccgo is slower to compile than the default gc, but supports more powerful
        optimizations so many programs compiled with it will run faster.
        This takes great care and even more patience.

    - The cgo program provides the mechanism for a foreign function interface to
        allow safe calling of C libraries from Go code. SWIG extends this capability
        to C++ libraries.
- Limitations of Go
    - No Generics, no .map, .reduce, .filter
    - No exceptions or assertions, as mentioned earlier (but I feel that this could go either way)
    - No Ternary operations. Use the good ole 'if else' statements.
    - Absolutely no tolerance for unused variables or imports (but is that really a bad thing?)
    - Your virus protection might think your compiled binaries are infected because it doesn't understand the structure of a Go Binary 😅
    - No pointer arithmetic (but thats for safety, and simplifies implementation of the garbage collector)
    - And honestly, dealing with GOPATH is kinda messy and annoying. You are forced to do all of your go projects within your GOPATH, but IDE's like VSCode and GoLand allow you to set the GoPath for your project without affecting your system's actual GOPATH.
-Philosophy
    - Do not communicate by sharing memory; instead, share memory by communicating
        The number 1 feature of GoLang is the exact opposite definition of what OOP stands for.
    - Keep It Simple Stupid (KISS)
        - Only one way of doing things
            No more spaces vs tabs, functions vs objects, for vs while …
            Developers write the same code, compiler can optimize it better.
- OOP
    - There are no classes in Go, but some features of Object Oriented programming are supported:
        - Encapsulation (possible on package level in Go)
        - Composition (possible through embedding in Go)
        - Polymorphism (possible through Interface satisfaction in Go)
        - Inheritance (Not possble - overloading makes code harder to debug)



# Why Golang 
----------
- Go is natively compiled
    - No VM needed. 
    - Directly compiled to machine code like c, c++.
- Static Typing 
- Scalable to large Systems 
- General Purpose Programming Language 
    - Mainly targeted at System/ Server-side programming like C, C++, Rust, ..
- Clean Syntax 
- Excellent Support for concurrency and networking
- Automatic Garbage Collection
- Rich Standard library
- Go Compiler is available on Linux, OS X, Windows, various BSD & Unix versions
- Go is open Source
		
# Who Uses Go 
-----------
- Many Google web properties and systems including youtube, Kubernetes containers and download server dl.google.com
- Docker, a set of tools for deploying Linux containers
- Dropbox, migrated some of their critical components from Python to Go
- SoundCloud, for many of their systems 
- CloudFoundry, aplatform as a service(PaaS)
- Couchbase, Query and indexing services within the couchbase server 
- MongoDB, tools for administering MongoDB instances 
- ThoughtWorks, some tools and applications around continous delivery and instant messages
- SendGrid, a transactional email delivery and management service 
- BBC, in some games and internal projects
- Novartis, for an internal inventory system

# What's not PRESENT in Go 
-----------------
- Class
- Class based Inheritance 
- Generics 
- Traits/ Mixins 
- Pointer Arithmetic
- Implicit type conversions 
- Method/Function Overloading

# Go Tools
- To check the go version,
        go version
- To know all the options,
        go help
- To get help about a package,
        godoc fmt Println

## DataTypes
    - statically typed programming language
    - Data Types
        - Integers
            - unsigned - uint8, uint16, uint32, uint64
            - signed - int8, int16, int32, int64
            - aliases:
                - rune is alias for int32
                - byte is alias for uint8
                    1 byte = 8 bits
                    1024 bytes = 1 kilobyte
                    1024 kilobytes = 1 Megabyte ...
            - Also, there are three machine dependent integer types
                - Machine dependent because their size depends on type of system architecture.
                - uint, int, uintptr
            - Generally, we can use int

        - Floating-point numbers
            - float32 - referred, also, as single precision
            - float64 - referred, also, as double precision
            - Larger sized floating-point number increases its precision
            - Generally, we can use float
        - NaN (not a number)
            - To represent 0/0
        - +∞ and −∞ (positive and negative infinity)
        - Complex
            - complex64
            - complex128

        - Strings
            - Go Strings are made up of individual bytes, usually one for each character
            - non-english characters are represented by more than 1 byte
            - created using
                - double quotes("Hello world")
                    - single-line strings
                    - Allow special sequences
                - back ticks (`Hello world`)
        - bool (true, false)

Identifier Naming Convention
    - first character - A-Z, a-z
    - remaining chars - A-Z, a-z, 0-9, _
    - pascalCase or CamelCase is recommended.

## Important Links:
    - tour.golang.org
    - lean.go.dev
    - gophercises.com
    - gobyexample.com
    - golangprograms.com
    - for jobs, golang.cafe
    - https://www.golang-book.com/
    - https://www.gopl.io/
    - https://www.goinggo.net/

    - https://github.com/GoesToEleven/svcc-19

    - Bill Kennedy
        - https://www.ardanlabs.com/blog/

    - https://www.twitter.com/joncalhoun?lang=en
    - https://www.gophercises.com/

    https://www.udemy.com/user/toddmcleod/
    https://www.greatercommons.com/

    https://www.changelog.com/gotime
    https://golangbyexample.com/

    https://golang.org/
    golang spec
    effective golang
    https://go.libhunt.com/
    forum.golangbridge.org

    https://programming.guide/go/
    godoc.org
        standard library and 3rd party packages

This source code is tested in go 1.13.5

## TODO:

- https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
- https://go101.org/article/channel.html
- https://exercism.io/tracks/go/learning

## Articles
    [Go CodeReview Comments guidelines](https://github.com/golang/go/wiki/CodeReviewComments)
    https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65
    [Farewell Node.js](https://medium.com/code-adventures/farewell-node-js-4ba9e7f3e52b?)

## Quotes

    “Go will be the server language of the future.” — Tobias Lütke, Shopify
    “With constant pressure to add features and options and configurations, and to
     ship code quickly, it’s easy to neglect simplicity, even though in the long
     run simplicity is the key to good software” — Rob Pike, One of Golang creators

## Packages

1. GoReleaser, http://github.com/goreleaser
2. Cobra( for creating CLIs), https://github.com/spf13/cobra
3. github.com/dgrijalva/jwt-go 				- generate and sign json web tokens
4. github.com/doug-martin/goqu 				- SQL query builder. Don't like ORMs in general, didn't like GORM. The API is a bit verbose but it does the job and supports tons of SQL features including database specific ones.
5. github.com/gin-gonic/gin 					- My goto framework for writing microservices
6. github.com/go-chi/chi 						- when net/http is not enough this is the router I use. gorilla/mux is slow, didn't bother trying anything else after chi.
7. github.com/go-gorm/gorm 					- database "orm"
8. github.com/golang-migrate/migrate/v4 		- database schema migration
9. github.com/golang/mock 						- Auto generate mocks from interfaces for unit testing
10. github.com/gonum/gonum 						- Contains almost everything one needs for numerical analysis. The guys over there are porting LAPACK to pure Go with explicit tests.
11. github.com/hashicorp/go-hclog 				- structured logging
12. github.com/jackc/pgx 						- excellent library for Postgre. Especially if you need PSQL specific features
13. github.com/kyleconroy/sqlc 					- generate funcs, structs and interfaces by providing a sql schema and queries
14. github.com/labstack/echo 					- web framework
15. github.com/nhooyr/websocket 				- websocket implementation. Recently migrated to it from gorilla/websocket. Nicer API and faster but nothing really wrong with gorilla. Just wanted to try something new.
16. github.com/pkg/errors 						- don't like how wrapping is done in standard library. This is better. Been using this library forever and it's even compatible with recent standard library changes.
17. github.com/rs/zerolog 						- structural logging library, fluent interface
18. github.com/satori/go.uuid  					- great for generating uuid's
19. github.com/shopspring/decimal  				- super useful when dealing with currency
20. github.com/sirupsen/logrus 					- go-to logger when i need something more than just pushing to stdout
21. github.com/stretchr/testify 				- very useful for tests
22. github.com/stretchr/testify#suite-package 	- I primarily use the suite package for a clear separation of setup and teardown during tests and the Require() to fail fast the test at the point of error
23. github.com/stretchr/testify/require 		- abort a test if some requirement is not met in one line
24. go.uber.org/automaxprocs 					- must-have for kubernetes deployments
25. pkg.go.dev/golang.org/x/sync/errgroup 		- Managing goroutines which can return error in the idomatic go way


# Upgrading dependencies
- To upgrade a specific package, 
    go get -u <package-name>

- To upgrade all (direct and indirect) packages in go.mod file,
    go get -u ./...

    For ony patch upgrade, (example  1.1.x)
    go get -u=patch ./...

    go get -u ./... walks all packages in your project. This is the command you want to use.
    go get -t -u ./... walks all packages in your project and also downloads tests files of these dependencies. Probably you don’t need that.
    go get -u updates in the current directory only. Useful for small single-package projects, just use the first version.
    go get -u specific.com/package updates just one (or more separated by space) packages (and dependencies).
    go get -u specific.com/package@version the same but to a specific version.
    go get -u all updates modules from the build list from go.mod. This is useful for listing (go list -m -u all) but not too useful for updates.