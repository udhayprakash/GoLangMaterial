Q1) What is the default value of a global variable in Go? Ans) A global variable has default value as its corresponding
zero-value.

Q2) What is the purpose of the function Printf()? Ans) Prints the formatted output

Q3) What is lvalue and rvalue? Ans) The expression appearing on right side of the assignment operator is called as
rvalue. Rvalue is assigned to Lvalue, which appears on left side of the assignment operator. Lvalue should designate to
a variable, not a constant.

Q4) What is the difference between actual & formal parameters? Ans) The parameters sent to the function at calling end
are called as actual parameters while at the receiving of the function definition called as formal parameters.

Q5) What is the difference between variable declaration & variable definition? Ans) Declaration associates type to the
variable whereas definition gives the value to the variable.

Q6) What is a token? Ans) A Go program consists of various tokens and a token is either a keyword, an identifier, a
constant, a string literal, or a symbol.

Q7) Which keyword is used to perform unconditional branching? Ans) goto

Q8) What is an array? Ans) Array is collection of similar data items under a common name.

Q9) What is structure in Go? Ans) Structure is another used defined data type available in Go Programming, which allows
you to combine dta items of different kinds.

Q10) How to define a structure in Go? Ans) To define a structure, you must use type and struct statememts. The struct
statement defines a new data type, with more than one member for your program. type statement binds a name with the type
which is struct in our case.

Syntax:
------

type struct_variable_type struct { member definition; member definition; ... member definition; }

Q11) What is slice in Go? Ans) Go Slice is an abstraction over Go Array. As Go Array allows you to define type of
variables that can hold several data items of the same kind but it do not provide any inbuilt method to increase size of
it dynamically or get a sub-array of its own. slices covers this limitation.

It provides many utility functions required on Array and is widely used in Go programming.

Q12) How to define a slice in Go? Ans) To define a slice, you can declare it as an array without specifying size or use
make function to crete the one.

var numbers []int /* a slice of unspecified size*/ numbers == []int{0, 0, 0, 0, 0, 0} numbers = make([]int, 5, 5) // a
slice of length 5 and capacity 5

Q13) How to get the count of elements present in a slice? Ans) len() function returns the elements present in the slice.

Q14) What is the difference between len() and cap() functions of slice in Go? Ans) len() function returns the elements
present in the slice whereas cap() function returns capacity of slice as how many elemnts it can be accommodate.

Q15) How to get a sub-slice of a slice? Ans) Slice allows lower-bound and upper-bound to be specified to get the
sub-slice of it using [lower_bound: upper_bound].

Q16) What is range in Go? Ans) The range keyword is used in for loop t iterate over items of an array, slice, channel or
map. With array and slices, it returns the index of the item as integer. With maps, it returns the key of the next
key-value pair.

Q17) What are maps in Go? Ans) Go provides another important data type map which maps unique keys to values. A key is an
object that you use to retrieve a value at a later date. Given a key and a value, you can store the value in a Map
object. After value is stored, you can retrieve it by using its key.

Q18) How to create a map in Go? Ans) You must use make function to create a map.

    // declare a variable, by default map will be nil
    var map_variable map[key_dta_type] value_data_type
    
    // define the map as nil. map can be assigned any value
    map_variable = make(map[key_data_type] value_data_type)

Q19) How to delete an entry from a map in Go? Ans) delete() function is used to delete an entry from the map. It
requires map and corresponding key which is to be deleted.

Q20) What is type casting in Go? Ans) Type casting is a way to convert a variable from one data type to another data
type. For example, if you want to store a long value into a simple integer then you can type cast long to int. You can
convert values from one type to another using the cast operator:
type_name(expression)

Q21) What are interfaces in Go? Ans) Go programming provides another data type called interfaces which represents a set
of method signatures. struct data type implements these interfaces to have method definitions for the method signature
of the interfaces.

Q22) In GO language how you can check variable type at runtime? Ans) A special type of switch is dedicated in GO to
check variable type at runtime, this switch is referred as type switch. Also, you can switch on the type of an interface
value with Type Switch.

Q23) Explain workspace in GO? Ans) Inside a workspace GO code must be kept. A workspace is a directory hierarchy with
three directories at its root. src contains GO source files organized into packages pkg contains package objects and bin
contains executable commands

Q24) How you can format a string without printing? Ans) To format a string without printing you have to use command
fmt.Sprintf( “at %v, %s”, e.When, e.What)

Q25) If a program is running across multiple threads, does failure in one goroutine halt the program? Ans) Yes, Eve
though running in multiple threads and cores, they are all part of same loop (coroutine).

Q26) Will running CPU intensive tasks in goroutines block them? Ans) May be. Because we are not sure whether all are
running on same core or not.

Q27) When to prefer concurrency, over Parallelism? Ans)    I/O Bound = Concurrency CPU Bound = Parallelism

https://programming.guide/go/go-gotcha.html