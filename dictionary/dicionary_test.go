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
		dictionary := Dictionary{}
		word := "rum"
		definition := "tasty alcoholic treat"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add word that already exists", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word:definition}
		err := dictionary.Add(word, "a new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update a word's definition", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word:definition}
		newDefinition := "a new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update definition when word does not exist", func(t *testing.T){
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)

	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
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
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}