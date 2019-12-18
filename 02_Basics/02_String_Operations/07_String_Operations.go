package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    // String comparision
    fmt.Println(strings.Compare("a", "b")) // -1
    fmt.Println(strings.Compare("a", "a")) //  0
    fmt.Println(strings.Compare("b", "a")) //  1

    // String Contains reports whether substr is within s.
    fmt.Println(strings.Contains("seafood", "foo")) // true
    fmt.Println(strings.Contains("seafood", "bar")) // false
    fmt.Println(strings.Contains("seafood", "")) // true
    fmt.Println(strings.Contains("", "")) // true

    // String ContainsAny reports whether any Unicode code points in chars are within s.
    fmt.Println(strings.ContainsAny("team", "i")) // false
    fmt.Println(strings.ContainsAny("fail", "ui")) // true
    fmt.Println(strings.ContainsAny("ure", "ui")) // true
    fmt.Println(strings.ContainsAny("failure", "ui")) // true
    fmt.Println(strings.ContainsAny("foo", "")) // false
    fmt.Println(strings.ContainsAny("", "")) // false

    // String ContainsRune reports whether the Unicode code point r is within s.
    // Finds whether a string contains a particular Unicode code point.
    // The code point for the lowercase letter "a", for example, is 97.
    fmt.Println(strings.ContainsRune("aardvark", 97)) // true
    fmt.Println(strings.ContainsRune("timeout", 97))  // false

    // String Count counts the number of non-overlapping instances of substr in s.
    // If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
    fmt.Println(strings.Count("cheese", "e"))  // 3
    fmt.Println(strings.Count("five", "")) // before & after each rune. 1 + 4 = 5

    // String EqualFold - case insensitive equivalence check
    fmt.Println(strings.EqualFold("Go", "go"))  // true
    fmt.Println(strings.EqualFold("Go", "GO"))  // true

    // String Fields splits string around one or more consecutive white space characters
    fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   ")) // ["foo" "bar" "baz"]
    fmt.Printf("Fields are: %q \n", strings.Fields("foo bar baz"))       // ["foo" "bar" "baz"]
    fmt.Printf("Fields are: %q \n", strings.Fields("foobarbaz"))         // ["foobarbaz"]

    // String FieldsFunc splits string around custom condition
    f := func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
    }
    fmt.Printf("Fields are: %q \n", strings.FieldsFunc("  foo1;bar2,baz3...", f)) // ["foo1" "bar2" "baz3"]

    // String HasPrefix tests whether the string s begins with prefix.
    fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
    fmt.Println(strings.HasPrefix("Gopher", "go")) // false
    fmt.Println(strings.HasPrefix("Gopher", ""))   // true

    // String HasSuffix tests whether the string s ends with suffix.
    fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
    fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
    fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
    fmt.Println(strings.HasSuffix("Amigo", ""))    // true

    // String Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
    fmt.Println(strings.Index("chicken", "c"))  // 0  -- first occurrence
    fmt.Println(strings.Index("chicken", "ken"))  // 4
    fmt.Println(strings.Index("chicken", "dmr"))  // -1

    // String IndexAny returns index of any character in substring

    fmt.Println(strings.IndexAny("chicken", "aeiouy")) // "i" is present  - 2
    fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1

    // String IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
    fmt.Println(strings.IndexByte("golang", 'g'))  // 0
    fmt.Println(strings.IndexByte("gophers", 'h')) // 3
    fmt.Println(strings.IndexByte("golang", 'x'))  // -1

    // String IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.
    f1 := func(c rune) bool {
        return unicode.Is(unicode.Han, c)
    }
    fmt.Println(strings.IndexFunc("Hello, 世界", f1))  // 7
    fmt.Println(strings.IndexFunc("Hello, world", f1)) // -1

    // String IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.
    fmt.Println(strings.IndexRune("chicken", 'k')) // 4
    fmt.Println(strings.IndexRune("chicken", 'd')) // -1

    // String concatenates joins strings to single string
    s := []string{"foo", "bar", "baz"}
    fmt.Println(strings.Join(s, ", "))  // foo, bar, baz
    fmt.Println(strings.Join(s, "__"))  // foo__bar__baz

    



}

