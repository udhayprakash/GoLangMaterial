# Structs

    - User-defined type that contains a collection of named fields/properties.
    - so, struct type variables are also called "named-types".
    - used to group related data together to form a single unit.
    - structs are as a lightweight class that supports composition but not inheritance.
    - Any struct type that starts with a capital letter is exported and accessible from outside packages. Similarly, any struct field that starts with a capital letter is exported.
    - On the contrary, all the names starting with a small letter are visible only inside the same package.

## copying structs

    - Structs are value-types
        - when a value-type instance is created, single space in memory is allocated
          to store a value.
    - when you assign one struct variable to another, a new copy of the struct is
        created and assigned.
    - Similarly, when you pass a struct to another function, the function gets its
        copy of the struct.
