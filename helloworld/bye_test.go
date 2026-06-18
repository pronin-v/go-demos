package helloworld

import (
	"fmt"
	"testing"
)

func TestBye(t *testing.T) {
	t.Run("Goodbye message for empty name", func(t *testing.T) {
		got := Bye("", "")
		want := "Bye bye, Dude"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Goodbye message for 'Kelly'", func(t *testing.T) {
		got := Bye("Kelly", "")
		want := "Bye bye, Kelly"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Goodby message for 'Spanish' lang", func(t *testing.T) {
		got := Bye("Diego", "Spanish")
		want := "Adios, Diego"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("Goodby message for 'Spanish' lang and empty name", func(t *testing.T) {
		got := Bye("", "Spanish")
		want := "Adios, amigo"
		AssertCorrectMessage(t, got, want)
	})
}

func ExampleBye() {
	msg := Bye("Pedro", "Spanish")
	fmt.Println(msg)
	// Output: Adios, Pedro
}
