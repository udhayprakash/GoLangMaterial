# arrays

    - Collection of elements of same data type
    - array length is fixed on declaration
        - size can neither be reduced or increased
        - For variable length arrays, go with "slice"
    - composite data type
        - composed of primitive data types like int, string, bool etc.

    - Arrays are values. Assigning one array to another copies all the elements.
    - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
    - The size of an array is part of its type. The types [10]int and [20]int are distinct.
