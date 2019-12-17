# GoLangMaterial
Go Language Material

    Go
        - General Purpose Compiled Programing language
        - No pre-requisites need. Can be learned as first programming language.
        - C language is the most direct influence on Go.
        - Platform-independent. 
            - Create executable build based on the platform. 
            - Means, source code is translated into language that the system understands.
        - comes with its own compiler
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
        - statically typed programming language

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
        - golangprograms.com
        - for jobs, golang.cafe
    
This source code is tested in go 1.13.5
