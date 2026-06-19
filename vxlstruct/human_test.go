package vxlstruct

import "testing"

func TestHuman(t *testing.T) {
	t.Run("Has ToString method", func(t *testing.T) {
		var h Human = Human{Gender: "Male", Age: 32}
		got := h.ToString()
		want := "gender: Male; age: 32;"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
