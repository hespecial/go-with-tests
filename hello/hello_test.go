package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("to a people", func(t *testing.T) {
		got := Hello("Hespecial", "")
		want := "Hello, Hespecial"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
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
