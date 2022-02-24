package main

import "fmt"

// Excersie to practise pointers. Swap variable values
func swap(x, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
