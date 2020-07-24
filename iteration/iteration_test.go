package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("test it repeates the characted 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("test it uses default count if counter is 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("test it uses default count if counter is -1", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
