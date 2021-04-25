# Methods

    - Functions with  a special reciever argument
    - The receiver argument has a name and a type.
        - It appears between the func keyword and the method name -
        - Syntax:
            func (receiver Type) MethodName(parameterList) (returnTypes) {
            }
    - receiver can be either a struct type or a non-struct type.
    - Methods help you avoid naming conflicts.
    - Since a method is tied to a particular receiver, you can have
      the same method names on different receiver types.
    - Methods in Go are more general than in C++ or Java: they can be defined
      for any sort of data, even built-in types such as plain, "unboxed" integers.
      They are not restricted to structs(classes).
