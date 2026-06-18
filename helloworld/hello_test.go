package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nick", "")
		want := "Hello, Nick"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' if empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lili", "French")
		want := "Bonjour, Lili"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Michail", "Russian")
		want := "Привет, Michail"
		AssertCorrectMessage(t, got, want)
	})
}
