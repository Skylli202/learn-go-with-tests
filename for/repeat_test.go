package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertExpected(t, expected, repeated)
	})
}

func assertExpected(t *testing.T, expected, received string) {
	if received != expected {
		t.Errorf("expected %q but got %q", expected, received)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

func ExampleRepeat() {
	repeated := Repeat("abab", 2)
	fmt.Printf("%s", repeated)
	// Output: abababab
}
