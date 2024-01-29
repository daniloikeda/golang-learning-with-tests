package slices

import ("testing"
	"slices")

func TestSumAll(t *testing.T) {
	listOfNumbers := SumAll([]int {1, 4}, []int{4, 7})
	want := []int {5, 11}

	if slices.Compare(listOfNumbers, want) != 0 {
		t.Errorf("Expected %v but got %v", want, listOfNumbers)
	}
}