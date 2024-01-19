package integers

import "testing"

func TestAdders(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		want := 4

		if sum != want {
			t.Errorf("expected %d want %d", sum, want)
		}
	})
}
