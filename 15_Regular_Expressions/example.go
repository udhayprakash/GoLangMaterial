package main

import (
	"fmt"
	"regexp"
)

// prettyMatches formats Matches nicely.
func prettyMatches(m []string) string {
	s := "["
	for i, e := range m {
		s += e
		if i < len(m)-1 {
			s += "|"
		}
	}
	s += "]"
	return s
}

// PrettySubmatches formats Submatches nicely.
func prettySubmatches(m [][]string) string {
	s := "[\n"
	for _, e := range m {
		s += "    " + prettyMatches(e) + "\n"
	}
	s += "]"
	return s
}

var (
	exps = []string{"b.*tter", "b(i|u)tter", `batter (\w+)`}

	text = `Betty Botter bought some butter 
But she said the butter’s bitter 
If I put it in my batter, it will make my batter bitter 
But a bit of better butter will make my batter better 
So ‘twas better Betty Botter bought a bit of better butter`
)

func main() {
	for _, e := range exps {
		re := regexp.MustCompile(e)
		fmt.Println(e + ":")
		fmt.Println("1. FindString: ", re.FindString(text))
		fmt.Println("2. FindStringIndex: ", re.FindStringIndex(text))
		fmt.Println("3. FindStringSubmatch: ", re.FindStringSubmatch(text))
		fmt.Printf("4. FindAllString: %v\n", prettyMatches(re.FindAllString(text, -1)))
		fmt.Printf("5. FindAllStringIndex: %v\n", re.FindAllStringIndex(text, -1))
		fmt.Printf("6. FindAllStringSubmatch: %v\n\n", prettySubmatches(re.FindAllStringSubmatch(text, -1)))
	}
}
