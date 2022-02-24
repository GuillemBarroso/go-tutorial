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

// Use NEW to get the pointer
func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)  // x is still 5
	zeroPointer(&x) // &variableName gives the address to "variableName" (pointer to the variable)
	fmt.Println(x)  // x is now 0

	xPtr_new := new(int) // new allocates memory for the type of variable and returns a pointer to it
	one(xPtr_new)
	fmt.Println(*xPtr_new) // x is 1
	/* GO is a garbage collected programming language, maning that memory is cleaned up automatically
	when nothing refers to it anymore.
	*/
}
