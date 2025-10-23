package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		// .2f -> f is flor float64, and 2 means print 2 decimal places
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			// '%g' will print a more precise decimal number in the error message
			t.Errorf("Expected %g, but got %g", want, got)
		}
	}

	t.Run("calculate area for rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rectangle, want)
		//got := rectangle.Area()
		//want := 72.0
		//
		//if got != want {
		//	t.Errorf("Expected %.2f, but got %.2f", want, got)
		//}
	})

	t.Run("calculate area for circles", func(t *testing.T) {
		circle := Circle{10}
		//got := circle.Area()
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func TestAreaTableTests(t *testing.T) {

	//  "anonymous struct" -> [] struct {}
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// go test -run TestArea/Rectangle
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				// %#v format string will print out our struct with the values in its field, so the developer
				// can see at a glance the properties that are being tested
				t.Errorf("[%#v] Expected %g, but got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
