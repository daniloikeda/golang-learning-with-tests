package integers

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("All positive numbers", func(t *testing.T) {
		orderedNums := BubbleSort([]int{3, 2, 1})
		want := []int{1, 2, 3}

		if !areEqual(t, orderedNums, want) {
			t.Errorf("expected %v want %v", orderedNums, want)
		}
	})
	t.Run("All negative numbers", func(t *testing.T) {
		orderedNums := BubbleSort([]int{-3, -2, -1})
		want := []int{-3, -2, -1}

		if !areEqual(t, orderedNums, want) {
			t.Errorf("expected %v want %v", orderedNums, want)
		}
	})
	t.Run("Mixed numbers", func(t *testing.T) {
		orderedNums := BubbleSort([]int{-3, 2, -1})
		want := []int{-3, -1, 2}

		if !areEqual(t, orderedNums, want) {
			t.Errorf("expected %v want %v", orderedNums, want)
		}
	})
}

func areEqual(t testing.TB, slice1 []int, slice2 []int) bool {
	t.Helper()

	if cap(slice1) != cap(slice2) || len(slice1) != len(slice2) {
		return false
	}

	for index, value := range slice1 {
		if value != slice2[index] {
			return false
		}
	}

	return true
}
