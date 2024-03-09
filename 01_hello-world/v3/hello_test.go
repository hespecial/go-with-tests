package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hespecial")
	want := "Hello, Hespecial"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
