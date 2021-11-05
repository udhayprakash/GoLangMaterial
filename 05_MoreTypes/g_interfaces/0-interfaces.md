In OOP, a protocol or interface is a common means for unrelated objects to communicate between themselves.

usage of interfaces:

    - writing generic algorithms
    - hiding implementation details
    - providing interception points

Go Types :
    1. abstract types - int8, int16, int32, int64, string, ... 
    2. concreate types - io.Reader, io.Writer, fmt.Stringer, ...

## union of interfaces

type ReadWriter interface { Reader Writer }

The bigger the interface, the weaker the abstraction - Rob Pike in his Go proverbs

## naming convention

Interfaces that specify just one method are usually just that function name with 'er' appended to it.

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

Sometimes the result isn't correct English, but we do it anyway:

    type Execer interface {
        Exec(query string, args []Value) (Result, error)
    }

Sometimes we use English to make it nicer:

    type ByteReader interface {
        ReadByte() (c byte, err error)
    }

When an interface includes multiple methods, choose a name that accurately describes its purpose (examples: net.Conn, http.ResponseWriter, io.ReadWriter).
 