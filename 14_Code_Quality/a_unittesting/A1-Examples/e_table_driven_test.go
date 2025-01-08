package main

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int64
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}

func TestOlder(t *testing.T) {
	// naked struct - creating and using immediately
	cases := []struct {
		person1  *Person
		person2  *Person
		expected bool
	}{
		{
			person1: &Person{
				name: "Ramesh",
				age:  22,
			},
			person2: &Person{
				name: "Prabas",
				age:  23,
			},
			expected: false,
		},
		{
			person1: &Person{
				name: "Suresh",
				age:  80,
			},
			person2: &Person{
				name: "Ramya",
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
	// naked struct - creating and using immediately
	cases := []TestCase2{
		{
			person1: &Person{
				name: "Rajanikanth",
				age:  78,
			},
			person2: &Person{
				name: "KamalHassan",
				age:  77,
			},
			expected: false,
		},
		{
			person1: &Person{
				name: "Modi",
				age:  80,
			},
			person2: &Person{
				name: "Cbn",
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

/*

$ go test -v e_table_driven_test.go 
=== RUN   TestOlder
    e_table_driven_test.go:55: Running: older(&{Ramya 80})
        Argument: &{Ramya 80} 
        Expected result: true
        Actual result: false
--- FAIL: TestOlder (0.00s)
=== RUN   TestOlder2
    e_table_driven_test.go:101: Running: older(&{KamalHassan 77})
        Argument: &{KamalHassan 77} 
        Expected result: false
        Actual result: true
--- FAIL: TestOlder2 (0.00s)
FAIL
FAIL    command-line-arguments  0.002s
FAIL

*/