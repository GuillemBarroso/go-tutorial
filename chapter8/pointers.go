package main

import "fmt"

// This function does not modify the value of x to be 0 because a copy of x is passed to the function zero(x)
func zero(x int) {
	x = 0
}

// To modify the input we need to point at it, use pointers
func zeroPointer(xPtr *int) { // *int means that is a pointer to an integer
	*xPtr = 0 // *variableName accesses the value of "variableName" (deferencing pointer; store 0 in the memory location xPtr refers to)
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)  // x is still 5
	zeroPointer(&x) // &variableName gives the address to "variableName"
	fmt.Println(x)  // x is now 0
}
