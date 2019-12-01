package services

import "testing"

func TestSort(t *testing.T) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	Sort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements) - 1] != 9 {
		t.Error("Last Element should be 9")
	}
}
