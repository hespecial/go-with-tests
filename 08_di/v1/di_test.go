package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hespecial")

	got := buffer.String()
	want := "Hello, Hespecial"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
