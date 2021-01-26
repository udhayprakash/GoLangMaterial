Slices
======

    - represent varaible-length sequences, []T
    - Looks like an array type without size
    - all elements are of same data type
    - slice has three components
        1. Pointer
            - Pointer points to first element of array that is reachable through slice,
              which is not necessarily the array's first element.
        2. Length
            - No. of slice elements; but cant exceed the capacity
        3. Capacity
            - No. of elements between the start of the slice and end of the underlying
              array.
              
    - Multiple slices can share the same underlying array and may refer to overlapping
      parts of that array.
      