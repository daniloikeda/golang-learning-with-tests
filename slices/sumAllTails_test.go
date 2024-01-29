package slices

import ("testing"
	"slices")

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func (t *testing.T) {

		got := SumAllTails([]int{2, 9, 15}, []int{5, 10, 25})
		expected := []int{24, 35}
	
		assertSlicesAreEqual(t, expected, got)
	})
	t.Run("safely sum empty slices", func (t *testing.T) {
		got := SumAllTails([]int{}, []int{1,2,3})
		expected := []int{0, 5}
	
		assertSlicesAreEqual(t, expected, got)
	})
}

func assertSlicesAreEqual(t testing.TB, expected, got []int) {
	t.Helper()
	if slices.Compare(got, expected) != 0 {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}