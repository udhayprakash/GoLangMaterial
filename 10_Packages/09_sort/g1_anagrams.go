package main

func main() {
	println(`AreAnagram("listen", "silent"):`, AreAnagram("listen", "silent"))
}

// AreAnagram - Anagram check solution with
// 		Time Complexity: O(N)
// 		Auxiliary Space: O(1)
func AreAnagram(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	counts := 0
	for _, runeValue := range word1 {
		counts += int(runeValue)
	}
	for _, runeValue := range word2 {
		counts -= int(runeValue)
	}
	println("counts:", counts)
	if counts != 0 {
		return false
	}
	return true
}
