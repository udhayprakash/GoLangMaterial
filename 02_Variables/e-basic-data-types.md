# Go Data Types

    1. Basic Types
        a) Numbers
            i) int
                1. signed  : int8, int16, int32, int64
                    - rune(represents a Unicode code point) is synonym for int32
                2. unsigned: uint8, uint16, uint32, uint64
                    - byte is synonym for uint8
                    - uintptr
                        - its width is both specified,
                           but is sufficient to hold all the bits of a pointer value
                        - Used, mainly, in low level programming
                3. general : int and uint - both either 32 or 64 bits

            ii) float(IEEE 754 standard implemented by all modern CPUs)
                - float32
                    - 6 decimal digits precision
                    - math.MaxFloat32 = 3.4e38
                    - math.MinFloat32 = 1.4e-45
                - float64 (recommended)
                    - 15 decimal digits precision
                    - math.MaxFloat64 = 1.8e308
                    - math.MinFloat64 = 4.9e-324
            iii) complex numbers
                    - complex64 complex128

        b) Strings
            - Strings
            - rune

        c) Booleans
            - true, false
    
        d) nil
            - zero value for pointers, interfaces, maps, slices, channels
              and function types, representing an uninitialized value.

    2. Aggregate Types
        - Values are concatenation of other values in memory
        - Fixed Size
        a) Arrays
            - Homogeneous - all elements are of same type
        b) Structs
            - Heterogeneous - elements may be of different type

    3. Reference Types
        - changes to them will be visible outside of the function
        a) Pointers
        b) Slices
        c) Maps
        d) Functions
        e) Channels

    4. Interface Types

Comparable Types: booleans, numbers, Strings

# NOTE:

- The int, uint, and uintptr types are usually 32 bits wide on
  32-bit systems and 64 bits wide on 64-bit systems.
- When you need an integer value you should use int unless you
  have a specific reason to use a sized or unsigned integer type.

# DataTypes

    ---------------------------------------------------------------------------------------------
    Data type	Description	                        Size	    Range
    ---------------------------------------------------------------------------------------------
    uint8	     8-bit unsigned integer	            1 byte	    0 to 255
    int8	     8-bit signed integer	            1 byte	    -128 to 127
    int16	    16-bit signed integer	            2 bytes	    -32768 to 32767
    unit16	    16-bit unsigned integer	            2 bytes	    0 to 65,535
    int32	    32-bit signed integer	            4 bytes	    -2,147,483,648 to 2,147,483,647
    uint32	    32-bit unsigned integer	            4 bytes	    0 to 4,294,967,295
    int64	    64-bit signed integer	            8 bytes	    -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
    uint64	    64-bit unsigned integer	            8 bytes	    0 to 18,446,744,073,709,551,615
    float32	    32-bit signed floating point number	4 bytes	    ±1.5e-45 to ±3.4e38
    float	    64-bit signed floating point number	8 bytes	    ±5.0e-324 to ±1.7e308
    string	    sequence of immutable text
    bool	    Stores either true or false	        1 byte	    True or false
    ---------------------------------------------------------------------------------------------
