package main

import (
	"fmt"
	"testing"
)

type Person struct {
	age  int64
	name string
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}
func TestOlder(t *testing.T) {
	cases := []struct {
		person1  *Person
		person2  *Person
		expected bool
	}{
		{
			person1: &Person{
				name: "Jane",
				age:  22,
			},
			person2: &Person{
				name: "John",
				age:  23,
			},
			expected: false,
		},
		{
			person1: &Person{
				name: "Michelle",
				age:  55,
			},
			person2: &Person{
				name: "Michael",
				age:  40,
			},
			expected: true,
		},
		{
			person1: &Person{
				name: "Ellen",
				age:  80,
			},
			person2: &Person{
				name: "Elliot",
				age:  80,
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := c.person1.older(c.person2)

		if actual != c.expected {
			out := fmt.Sprintf("Running: older(%v)\n", c.person2) +
				fmt.Sprintf("Argument: %v \n", c.person2) +
				fmt.Sprintf("Expected result: %t\n", c.expected) +
				fmt.Sprintf("Actual result: %t\n", actual)
			t.Fatalf(out)
		}
	}
}

// writing struct separately
type TestCase2 struct {
	person1  *Person
	person2  *Person
	expected bool
}

func TestOlder2(t *testing.T) {
	cases := []TestCase2{
		{
			person1: &Person{
				name: "Ellen",
				age:  80,
			},
			person2: &Person{
				name: "Elliot",
				age:  80,
			},
			expected: true,
		},
	}

	for _, c := range cases {
		actual := c.person1.older(c.person2)

		if actual != c.expected {
			out := fmt.Sprint("Argument: ", c.person2) +
				fmt.Sprintf("Expected result: %t\n", c.expected) +
				fmt.Sprintf("Actual result: %t\n", actual)
			t.Fatalf(out)
		}
	}
}

/*
-go test -v e_table_driven_test.go                                                   === RUN   TestOlder
    e_table_driven_test.go:65: Running: older(&{80 Elliot})
        Argument: &{80 Elliot}
        Expected result: true
        Actual result: false
--- FAIL: TestOlder (0.00s)
=== RUN   TestOlder2
    e_table_driven_test.go:99: Argument: &{80 Elliot}Expected result: true
        Actual result: false
--- FAIL: TestOlder2 (0.00s)
FAIL
FAIL    command-line-arguments  0.120s
FAIL
*/
