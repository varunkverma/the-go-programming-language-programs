package main

// Point represents a 2-D point
type Point struct {
	X, Y int
}

// Circle represents a 2-D circle
type Circle struct {
	Center Point
	Radius int
}

// Wheel represents a 2-D wheel
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}
