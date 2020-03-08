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
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 19.0, Height: 5.0}, hasArea: 95},
		{name: "Circle", shape: Circle{Radius: 15}, hasArea: 47.12388980384689},
		{name: "Triangle", shape: Triangle{Width: 43.0, Height: 5.0}, hasArea: 107.5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
