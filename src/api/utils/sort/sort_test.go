package sort

import "testing"

func TestBubbleSortOrderASC(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last Element should be 9")
	}
}

func TestBubbleSortASC(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("First Element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("Last Element should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}

}

func BenchmarkSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}

}
