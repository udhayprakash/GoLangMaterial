package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("expected %d; got %d", 3, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(3, 2)
	if result != 1 {
		t.Errorf("expected %d; got %d", 1, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(1, 2)
	if result != 2 {
		t.Errorf("expected %d; got %d", 2, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(2, 1)
	if result != 2 {
		t.Errorf("expected %d; got %d", 2, result)
	}
}
