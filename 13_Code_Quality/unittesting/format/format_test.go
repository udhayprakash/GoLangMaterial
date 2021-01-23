package format

import "testing"

func TestPrint(t *testing.T) {
	result := Print(3, 5)
	if result != "3.00000" {
		t.Errorf("expected %s; got %s", "3.00000", result)
	}
}
