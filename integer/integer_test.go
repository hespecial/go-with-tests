package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(7, 8)
	expected := 15
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
