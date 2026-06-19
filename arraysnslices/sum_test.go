package arraysnslices

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 4})
	want := []int{6, 8}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3, 4, 5, 6}, []int{9, 8, 7, 6})
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {

		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 4})
		want := []int{5, 4}
		checkSums(t, got, want)
	})

	t.Run("safe sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 4})
		want := []int{0, 4}
		checkSums(t, got, want)
	})
}
