package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("hespecial", "")
		want := "Hello hespecial"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("hespecial", "Spanish")
		want := "Hola hespecial"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("hespecial", "French")
		want := "Bonjour hespecial"
		assertCorrectMessage(t, got, want)
	})
}
