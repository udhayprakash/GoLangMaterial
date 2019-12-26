# GoLangMaterial
Go Language Material

    Go
        - General Purpose Compiled Programing language
        - created by Ken Thompson (B, C, Unix, UTF-8), Rob Pike(Unix, UTF-8), 
          Robert Griesemer(Hotspot, JVM) and few google engineers.
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

    Go Tools
        - To check the go version, 
                go version
        - To know all the options, 
                go help
        - To get help about a package, 
                godoc fmt Println

    DataTypes
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

Important Links:
----------------
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
        
        
        https://golang.org/
        golang spec
        effective golang
        https://go.libhunt.com/
        forum.golangbridge.org
        
        https://programming.guide/go/
        godoc.org
            standard library and 3rd party packages
This source code is tested in go 1.13.5


TODO:
-----
    - https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
    - https://go101.org/article/channel.html
    - https://exercism.io/tracks/go/learning

Articles
--------
    https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65
    [Farewell Node.js](https://medium.com/code-adventures/farewell-node-js-4ba9e7f3e52b?)
Quotes
------
    “Go will be the server language of the future.” — Tobias Lütke, Shopify 
    “With constant pressure to add features and options and configurations, and to 
     ship code quickly, it’s easy to neglect simplicity, even though in the long 
     run simplicity is the key to good software” — Rob Pike, One of Golang creators 
    