package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMesage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Luke", "")
		want := "Hello, Luke"
		assertCorrectMesage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMesage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodi", "Spanish")
		want := "Hola, Elodi"
		assertCorrectMesage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Blaze", "French")
		want := "Bonjour, Blaze"
		assertCorrectMesage(t, got, want)
	})
}
