package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("peremeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %f", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{19.0, 5.0}
		checkArea(t, rectangle, 95.0)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{15}
		checkArea(t, circle, 47.12388980384689)
	})

}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 19.0, Height: 5.0}, want: 95},
		{shape: Circle{Radius: 15}, want: 47.12388980384689},
		{shape: Triangle{Width: 43.0, Height: 5.0}, want: 107.5},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
