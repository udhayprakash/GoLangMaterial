Keywords (24)
--------
break       case        chan            const   continue    default
defer       else        fallthrough     for     func        go
if          import      interface       iota    map         nil
range       return      struct          switch  type        var


Built-in Functions (11)
------------------
append      cap     copy    delete  len
make        new     panic   print   println
recover

Operators
---------
    Arithmetic
    ----------
    
        -------------------------------------------------
        Operator	Name	            Types
        -------------------------------------------------
        +	        sum	                integers, floats, complex values, strings
        -	        difference	        integers, floats, complex values
        *	        product	            integers, floats, complex values
        /	        quotient	        integers, floats, complex values
        %	        remainder	        integers
        &	        bitwise AND	        integers
        |	        bitwise OR	        integers
        ^	        bitwise XOR	        integers
        &^	        bit clear (AND NOT)	integers
        <<	        left shift	        integer << unsigned integer
        >>	        right shift	        integer >> unsigned integer

    Comparison
    ----------
        ----------------------------------------------------------------------
        Operator	Name	            Types
        ----------------------------------------------------------------------
        ==	        equal	            comparable
        !=	        not equal	        comparable
        <	        less	            integers, floats, strings
        <=	        less or equal	    integers, floats, strings
        >	        greater	            integers, floats, strings
        >=	        greater or equal	integers, floats, strings
        ----------------------------------------------------------------------

        - Boolean, integer, floats, complex values and strings are comparable.
        - Strings are ordered lexically byte-wise.
        - Two pointers are equal if they point to the same variable or if both are nil.
        - Two channel values are equal if they were created by the same call to make or if both are nil.
        - Two interface values are equal if they have identical dynamic types and equal concrete values or if both are nil.
        - A value x of non-interface type X and a value t of interface type T are equal if t's dynamic type is identical to X and t's concrete value is equal to x.
        - Two struct values are equal if their corresponding non-blank fields are equal.
        - Two array values are equal if their corresponding elements are equal.
      
    Logical 
    -------
    
        --------------------------------------------------------------------------
        Operator	Name	            Description
        --------------------------------------------------------------------------
        &&	        conditional AND	    p && q denotes "if p then q else false"
        ||	        conditional OR	    p || q denotes "if p then true else q"
        !	        NOT	!p denotes      "not p"  
        --------------------------------------------------------------------------

    Others
    -------

        --------------------------------------------------------------------------
        Operator	Name	            Description
        --------------------------------------------------------------------------
        &	        address of	        &x generates a pointer to x
        *	        pointer indirection	*x denotes the variable pointed to by x
        <-	        receive	            <-ch is the value received from channel ch
        ...         Spread operator     To unpack the values in a string/array, ...
        --------------------------------------------------------------------------

Packages
--------
bufio   
bytes
compress/gzip
container/list
encoding/base64
encoding/csv
encoding/json
flag
fmt
html/template
io/ioutil
index/suffixarray
math
math/bits
net/http
os
os/exec
path
path/filepath
regexp
sort
strconv
rand
strings
time
unicode


// Ref: https://golang.org/pkg/builtin/