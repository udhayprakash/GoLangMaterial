In OOP, a protocol or interface is a common means for unrelated objects
to communicate between themselves.

usage of interfaces:
- writing generic algorithms
- hiding implementation details
- providing interception points

Go Types 
	1. abstract types - int8, int16, int32, int64, string, ...
	2. concreate types - io.Reader, io.Writer, fmt.Stringer, ...

union of interfaces
-------------------
type ReadWriter interface {
	Reader
	Writer
}

The bigger the interface, the weaker the abstraction - Rob Pike in his Go proverbs
