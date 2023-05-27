package math

import (
	"testing"
)

func TestAddore(t *testing.T) {
	result := AddCore(2.0, 3.0)
	if result != 5.0 {
		t.Errorf("Expected 5.0, but got %v", result)
	}
}

func TestSumCore(t *testing.T) {
	result := SumCore([]float64{1.0, 2.0, 3.0})
	if result != 6.0 {
		t.Errorf("Expected 6.0, but got %v", result)
	}
}
