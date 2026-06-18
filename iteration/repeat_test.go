package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Print 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("repeated %q expected %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}
