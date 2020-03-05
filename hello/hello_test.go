package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Luke")
	want := "Hello, Luke"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
