package anagram

import (
	"testing"
)

func TestAreAnagram(t *testing.T) {
	if AreAnagram("listen", "silent") != true {
		t.Error(`"listen", "silent"`)
	}
	if AreAnagram("test", "ttew") != false {
		t.Error(`"test", "ttew"`)
	}
	if AreAnagram("geeksforgeeks", "forgeeksgeeks") != true {
		t.Error(`"geeksforgeeks", "forgeeksgeeks"`)
	}
	if AreAnagram("triangle", "integral") != true {
		t.Error(`"triangle", "integral"`)
	}
	if AreAnagram("abd", "acb") != false {
		t.Error(`"abd", "acb"`)
	}
}
