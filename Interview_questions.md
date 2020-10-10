
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

**Ans)** Every GO program is made up of packages.  The program starts running in package main.  This program is using the packages with import paths “fmt” and “math/rand”.

**5)** Explain workspace in GO?

**Ans)** Inside a workspace GO code must be kept.  A workspace is a directory hierarchy with three directories at its root.

        src contains GO source files organized into packages
            pkg contains package objects and
            bin contains executable commands