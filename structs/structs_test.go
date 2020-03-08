package main

import "testing"

type Rectangle struct {
	width  float64
	height float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{19.0, 5.0}
	got := Area(rectangle)
	want := 95.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
