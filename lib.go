package main

import "testing"

const (
	spanish = "Spanish"
	french  = "French"
	russian = "Russian"
)

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
