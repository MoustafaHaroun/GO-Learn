package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Moustafa", "")
		want := "Hello, Moustafa"

		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		asserCorrectMessage(t, got, want)
	})

	t.Run("spanish test", func(t *testing.T) {
		got := Hello("Moustafa", "Spanish")
		want := "Hola, Moustafa"

		asserCorrectMessage(t, got, want)
	})

	t.Run("french test", func(t *testing.T) {
		got := Hello("Moustafa", "French")
		want := "Bonjour, Moustafa"

		asserCorrectMessage(t, got, want)
	})

	t.Run("ducth test", func(t *testing.T) {
		got := Hello("Moustafa", "Dutch")
		want := "Hallo, Moustafa"

		asserCorrectMessage(t, got, want)
	})
}

func asserCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q  want %q", got, want)
	}
}
