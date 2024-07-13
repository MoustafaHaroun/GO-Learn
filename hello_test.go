package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Moustafa")
	want := "Hello, Moustafa!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
