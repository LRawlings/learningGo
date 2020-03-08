package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test":"this is just a test"}

	t.Run("search for word that exist in dictionary", func(t *testing.T){
		got, _ := dictionary.Search("test")
		want := "this is just a test"
	
		assertStrings(t, got, want)
	})

	t.Run("search for word that doesn't exist in dictionary", func(t *testing.T){
		_, got := dictionary.Search("rum")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T){
	t.Run("add word and definition to dictionary", func(t *testing.T){
		dictionary := Dictionary{"test":"this is just a test"}
		dictionary.Add("rum", "tasty alcoholic treat")

		want := "tasty alcoholic treat"
		got, err := dictionary.Search("rum")
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}