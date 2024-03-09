package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Hespecial", "")
		want := "Hello, Hespecial"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Hespecial", spanish)
		want := "Hola, Hespecial"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Hespecial", french)
		want := "Bonjour, Hespecial"
		assertCorrectMessage(t, got, want)
	})

}
