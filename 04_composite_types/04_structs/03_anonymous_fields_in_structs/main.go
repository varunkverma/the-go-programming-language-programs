package main

import "fmt"

// Point represents a 2-D point
type Point struct {
	X, Y int
}

// Circle represents a 2-D circle
type Circle struct {
	Point  // anonymous field (must be a named type)
	Radius int
}

// Wheel represents a 2-D wheel
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Point.X = 8
	//w.Circle.Point.Y = 8
	w.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	// using struct literal
	w2 := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%v\n", w2)  // {{{8 8} 5} 20}
	fmt.Printf("%#v\n", w2) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w3 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 10,
				Y: 10,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%v\n", w3)
	fmt.Printf("%#v\n", w3)
}
