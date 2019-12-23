Methods
=======
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
