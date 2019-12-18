package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// String comparision
	fmt.Println("\n ==== strings.Compare  ====")
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) //  0
	fmt.Println(strings.Compare("b", "a")) //  1

	// String Contains reports whether substr is within s.
	fmt.Println("\n ==== strings.Contains ====")
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true

	// String ContainsAny reports whether any Unicode code points in chars are within s.
	fmt.Println("\n ==== strings.ContainsAny ====")
	fmt.Println(strings.ContainsAny("team", "i"))     // false
	fmt.Println(strings.ContainsAny("fail", "ui"))    // true
	fmt.Println(strings.ContainsAny("ure", "ui"))     // true
	fmt.Println(strings.ContainsAny("failure", "ui")) // true
	fmt.Println(strings.ContainsAny("foo", ""))       // false
	fmt.Println(strings.ContainsAny("", ""))          // false

	// String ContainsRune reports whether the Unicode code point r is within s.
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println("\n ==== strings.ContainsRune ====")
	fmt.Println(strings.ContainsRune("aardvark", 97)) // true
	fmt.Println(strings.ContainsRune("timeout", 97))  // false

	// String Count counts the number of non-overlapping instances of substr in s.
	// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
	fmt.Println("\n ==== strings.Count ====")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune. 1 + 4 = 5

	// String EqualFold - case insensitive equivalence check
	fmt.Println("\n ==== strings.EqualFold ====")
	fmt.Println(strings.EqualFold("Go", "go")) // true
	fmt.Println(strings.EqualFold("Go", "GO")) // true

	// String Fields splits string around one or more consecutive white space characters
	fmt.Println("\n ==== strings.Fields ====")
	fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   ")) // ["foo" "bar" "baz"]
	fmt.Printf("Fields are: %q \n", strings.Fields("foo bar baz"))       // ["foo" "bar" "baz"]
	fmt.Printf("Fields are: %q \n", strings.Fields("foobarbaz"))         // ["foobarbaz"]

	// String FieldsFunc splits string around custom condition
	fmt.Println("\n ==== strings.FieldsFunc ====")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q \n", strings.FieldsFunc("  foo1;bar2,baz3...", f)) // ["foo1" "bar2" "baz3"]

	// String HasPrefix tests whether the string s begins with prefix.
	fmt.Println("\n ==== strings.HasPrefix ====")
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "go")) // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true

	// String HasSuffix tests whether the string s ends with suffix.
	fmt.Println("\n ==== strings.HasSuffix ====")
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
	fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
	fmt.Println(strings.HasSuffix("Amigo", ""))    // true

	// String Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	fmt.Println("\n ==== strings.Index ====")
	fmt.Println(strings.Index("chicken", "c"))   // 0  -- first occurrence
	fmt.Println(strings.Index("chicken", "ken")) // 4
	fmt.Println(strings.Index("chicken", "dmr")) // -1
	fmt.Println()

	// String IndexAny returns index of any character in substring
	fmt.Println("\n ==== strings.IndexAny ====")
	fmt.Println(strings.IndexAny("chicken", "aeiouy")) // "i" is present  - 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))   // -1

	// String IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
	fmt.Println("\n ==== strings.IndexByte ====")
	fmt.Println(strings.IndexByte("golang", 'g'))  // 0
	fmt.Println(strings.IndexByte("gophers", 'h')) // 3
	fmt.Println(strings.IndexByte("golang", 'x'))  // -1

	// String IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.
	fmt.Println("\n ==== strings.IndexFunc ====")
	f1 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f1))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f1)) // -1

	// String IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.
	fmt.Println("\n ==== strings.IndexRune ====")
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1

	// String concatenates joins strings to single string
	fmt.Println("\n ==== strings.Join ====")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // foo, bar, baz
	fmt.Println(strings.Join(s, "__")) // foo__bar__baz

	// String LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
	fmt.Println("\n ==== strings.LastIndex ====")
	fmt.Println(strings.Index("go gopher", "go"))         // 0
	fmt.Println(strings.LastIndex("go gopher", "go"))     // 3
	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1

	// String LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
	fmt.Println("\n ==== strings.LastIndexAny  ====")
	fmt.Println(strings.LastIndexAny("go gopher", "go"))     // 4
	fmt.Println(strings.LastIndexAny("go gopher", "rodent")) // 8
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))   // -1

	// String LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
	fmt.Println("\n ==== strings.LastIndexByte  ====")
	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) // 10
	fmt.Println(strings.LastIndexByte("Hello, world", 'o')) //  8
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) // -1

	// String LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.
    fmt.Println("\n ==== strings.LastIndexFunc  ====")
    fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) //  5
    fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) //  2
    fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))     // -1

    // String Map
    fmt.Println("\n ==== strings.Map  ====")
    rot13 := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return 'A' + (r-'A'+13)%26
        case r >= 'a' && r <= 'z':
            return 'a' + (r-'a'+13)%26
        }
        return r
    }
    fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
    // 'Gjnf oevyyvt naq gur fyvgul tbcure...

    // String Repeat returns a new string consisting of count copies of the string s.
    //        It panics if count is negative or if the result of (len(s) * count) overflows.
    fmt.Println("\n ==== strings.Repeat  ====")
    fmt.Println("ba" + strings.Repeat("na", 2)) // banana

    // String func Replace(s, old, new string, n int) string
    //  If n < 0, there is no limit on the number of replacements.
    fmt.Println("\n ==== strings.Replace  ====")
    fmt.Println(strings.Replace("10001000101", "6", "33", 2))  // 10001000101
    fmt.Println(strings.Replace("10001000101", "1", "33", 2))  // 3300033000101
    fmt.Println(strings.Replace("10001000101", "1", "33", -1)) // 330003300033033

    // String ReplaceAll(s, old, new string) string
    fmt.Println("\n ==== strings.ReplaceAll  ====")
    fmt.Println(strings.ReplaceAll("10001000101", "6", "33"))  // 10001000101
    fmt.Println(strings.ReplaceAll("10001000101", "1", "33"))  // 3300033000101

    // String func Split(s, sep string) []string
    fmt.Println("\n ==== strings.Split  ====")
    fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // ["a" "b" "c"]
    fmt.Printf("%q\n", strings.Split("a,,b,c", ",")) // ["a" "" "b" "c"]

    fmt.Printf("%q\n", strings.Split("a b c", " ")) // ["a" "b" "c"]
    fmt.Printf("%q\n", strings.Split("a  b  c", " ")) // ["a" "" "b" "" "c"]

    fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
    // ["" "man " "plan " "canal panama"]

    fmt.Printf("%q\n", strings.Split(" xyz ", ""))  // [" " "x" "y" "z" " "]
    fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins")) // [""]

    // String func SplitAfter(s, sep string) []string
    // SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
    fmt.Println("\n ==== strings.SplitAfter  ====")
    fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
    fmt.Printf("%q\n", strings.SplitAfter("a,,b,c", ","))

    // func SplitAfterN(s, sep string, n int) []string
    // n > 0 : at most n substrings; the last substring will be the unsplit remainder.
    // n == 0: the result is nil (zero substrings)
    // n < 0 : all substrings
    fmt.Println("\n ==== strings.SplitAfterN  ====")
    fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
    fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 0)) // []
    fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", -1))// ["a," "b," "c"]

    // func SplitN(s, sep string, n int) []string
    // n > 0: at most n substrings; the last substring will be the unsplit remainder.
    // n == 0: the result is nil (zero substrings)
    // n < 0: all substrings
    fmt.Println("\n ==== strings.SplitN  ====")
    fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
    z := strings.SplitN("a,b,c", ",", 0)
    fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)

    // string  Title
    fmt.Println("\n ==== strings.Title  ====")
    // Compare this example to the ToTitle example.
    fmt.Println(strings.Title("her royal highness"))
    fmt.Println(strings.Title("loud noises"))
    fmt.Println(strings.Title("хлеб"))

    // string ToTitle
    fmt.Println("\n ==== strings.ToTitle  ====")
    // Compare this example to the Title example.
    fmt.Println(strings.ToTitle("her royal highness"))
    fmt.Println(strings.ToTitle("loud noises"))
    fmt.Println(strings.ToTitle("хлеб"))

    fmt.Println("\n ==== strings.ToTitleSpecial  ====")
    fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

    // string  ToLower returns s with all Unicode letters mapped to their lower case.
    fmt.Println("\n ==== strings.ToLower  ====")
    fmt.Println(strings.ToLower("Gopher"))

    // func ToLowerSpecial(c unicode.SpecialCase, s string) string
    fmt.Println("\n ==== strings.ToLowerSpecial  ====")
    fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

    // func ToUpper(s string) string
    fmt.Println("\n ==== strings.ToUpper  ====")
    fmt.Println(strings.ToUpper("Gopher is good"))


    // func ToUpperSpecial(c unicode.SpecialCase, s string) string
    fmt.Println("\n ==== strings.ToUpperSpecial  ====")
    fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

    // func Trim(s string, cutset string) string
    fmt.Println("\n ==== strings.Trim  ====")
    fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers
    fmt.Println(strings.Trim(" Hello, Gophers  ", " "))     // Hello, Gophers
    fmt.Println(strings.Trim(" Hello, Gophers  ", ""))      //   Hello, Gophers

    // func TrimFunc(s string, f func(rune) bool) string
    fmt.Println("\n ==== strings.TrimFunc  ====")
    fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    })) // Hello, Gophers

    // func TrimLeft(s string, cutset string) string
    fmt.Println("\n ==== strings.TrimLeft  ====")
    fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")) // Hello, Gophers!!!


    // func TrimLeftFunc(s string, f func(rune) bool) string
    fmt.Println("\n ==== strings.TrimLeftFunc  ====")
    fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    }))  // Hello, Gophers!!!

    // func TrimPrefix(s, prefix string) string
    fmt.Println("\n ==== strings.TrimPrefix  ====")
    var s1 = "¡¡¡Hello, Gophers!!!"
    s1 = strings.TrimPrefix(s1, "¡¡¡Hello, ")
    s1 = strings.TrimPrefix(s1, "¡¡¡Howdy, ")
    fmt.Print(s1)  // Gophers!!!

    // func TrimRight(s string, cutset string) string
    fmt.Println("\n ==== strings.TrimRight  ====")
    fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
    // ¡¡¡Hello, Gophers

    // func TrimRightFunc(s string, f func(rune) bool) string
    fmt.Println("\n ==== strings.TrimRightFunc  ====")
    fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r)
    })) // ¡¡¡Hello, Gophers

    // func TrimSpace(s string) string
    fmt.Println("\n ==== strings.TrimSpace  ====")
    fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))

    // func TrimSuffix(s, suffix string) string
    fmt.Println("\n ==== strings.TrimSuffix  ====")
    var s2 = "¡¡¡Hello, Gophers!!!"
    s2 = strings.TrimSuffix(s2, ", Gophers!!!")
    s2 = strings.TrimSuffix(s2, ", Marmots!!!")
    fmt.Print(s2)














}
