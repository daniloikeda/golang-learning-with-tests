package main

import ("testing")

func TestDisplacement(t *testing.T) {
	t.Run("Calculate displacement after 0 seconds", func (t *testing.T) {
		got := GenDisplaceFn(10, 2, 1)
		want := 1.0

		assertCorrectDisplacement(t, got(0.0), want)
	})
	t.Run("Calculate displacement after 3 seconds", func (t *testing.T) {
		got := GenDisplaceFn(10, 2, 1)
		want := 52.0

		assertCorrectDisplacement(t, got(3.0), want)
	})
}

func assertCorrectDisplacement(t testing.TB, got float64, want float64) {
	t.Helper();

	if (got != want) {
		t.Errorf("Expected %f and got %f", want, got)
	}
}