package main

import (
	"fmt"
	"math"
)

/* Improve struct.go by using methods
(c *Circle) is the receiver so the func area is callable
like c.area()
*/
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectagle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectagle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())

	r := Rectagle{0, 0, 10, 10}
	fmt.Println(r.area())
}
