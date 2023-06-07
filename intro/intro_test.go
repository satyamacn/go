package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hw("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hw("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hw("Elen", spanish)
		want := "Hola, Elen"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hw("Elen", french)
		want := "Bonjour, Elen"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	//t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
