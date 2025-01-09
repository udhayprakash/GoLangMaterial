## Fuzz testing (fuzzing)

It is a technique for testing software by providing it with random or semi-random inputs in order to find vulnerabilities and bugs. 

In Go, for fuzz testing
    before 1.18, https://github.com/dvyukov/go-fuzz is used. 
    From 1.18, supported with builtin package


types of Fuzzers 
===============
1) By Knowledge 
    - Black Box 
    - White Box 
    - Grey Box 
2) By input sample generation
    - Random
    - Mutation-Based 
    - Structure-aware
    - Behavior-aware

3) By Domain
    - Universal
    - Web Application
    - Embedded Systems
    - etc


Ref: https://python.plainenglish.io/friendly-fuzzing-tests-in-python-intro-72201778b1b5
