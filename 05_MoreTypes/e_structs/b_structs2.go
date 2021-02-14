package main

type Thing struct {
	Foo string
	Bar int
	Baz float64
}

func main() {
	m := map[string]Thing{
		"a": Thing{Foo: "abc", Bar: 7, Baz: 0.123},
		"b": Thing{Foo: "Hey", Bar: 0, Baz: 0.999},
		"c": Thing{Foo: "wat", Bar: 2, Baz: 1.234},
	}

	for k, v := range m {
		if v.Bar%2 == 0 {
			println(k, "has an even Bar")
		} else {
			println(k, "has an odd Bar")
		}
	}
}
