package main

import (
	"fmt"
	"math"
)

/* Only using Go's build-in data tpyes. Possible but can become quite tedious
and can affect the code's readability

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}
func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}
func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
}
*/

// Define a struct (contains name fields)
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// Initialise
	// var c Circle
	c := Circle{x: 0, y: 0, r: 5}
	//c := Circle{0, 0, 5}
	// c := new(Circle)

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(&c))
}
