## Design Patterns

    - Design patterns identify solutions to common programming problems, but they are
      not actual code ot algorithms themselves
    - Patterns are a form of best practices for approaching a problem
    - Usage of patterns can be significantly influenced by the language.
        - Ex: patterns useful in OOP programming may not work in non-OOP languages.

### Go Language features that affect design patterns

    - No support for traditional OOP classes, like C++, Java, C#
        - No static members or constructors - affects patterns like Singleton
    - No support for clas-based inhereitance or method overloading
        - Affects patterns like Visitor
    - No support for exceptions - error handling is via return values
        - Affects patterns like Builder
    - No support for abstract classes
        - Affects patterns like Abstract Factory and Bridge

### Creational Design Patterns

    - Deals with how objects are created during the life-time of the program.
    - Types
        - Abstract Factory
        - Builder
        - Factory
        - Lazy Initialization
        - Multiton
        - Object Pool
        - Prototype
        - Singleton

### Behavioural Design Patterns

    - Deals with organizing program structures
    - Types
        - Chain of Responsibility
        - Command
        - Interpreter
        - Iterator
        - Mediator
        - Memento
        - Null Object
        - Observer
        - State
        - Strategy
        - Template Method
        - Visitor

### Structural Design Patterns

    - Deals with increasing scale and communication
    - Types
        - Adapater
        - Bridge
        - Composite
        - Decorator
        - Facade
        - Flyweight
        - Proxy


ref: https://golangbyexample.com/all-design-patterns-golang/
