package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("1 + 1 = 2", func(t *testing.T) {
		sum := Add(1, 1)
		expected := 2
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4
}
