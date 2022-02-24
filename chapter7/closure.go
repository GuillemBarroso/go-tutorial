package main

import "fmt"

func main() {
	x := 0
	increment := func() int { // nested function accesing non-local variable = closure
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
