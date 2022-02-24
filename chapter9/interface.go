package main

import (
	"fmt"
	"math"
)

/* Interface is like a method but with a method set instead of defining fields
 */

type Shape interface {
	area() float64
	perimeter() float64
}

func totalQuantities(shapes ...Shape) (float64, float64) {
	var area float64
	var perimeter float64
	for _, s := range shapes {
		area += s.area()
		perimeter += s.perimeter()
	}
	return area, perimeter
}

type Circle struct {
	x, y, r float64
}

type Rectagle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectagle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectagle) perimeter() float64 {
	return 2 * (distance(r.x1, r.y1, r.x1, r.y2) * distance(r.x1, r.y1, r.x2, r.y1))
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectagle{0, 0, 10, 10}
	fmt.Println(totalQuantities(&c, &r))
}
