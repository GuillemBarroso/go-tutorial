package main

import "fmt"

// exercise to practise recursive functions. Fibonacci sequence
func fibonacci(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2) // much simpler!
}

func main() {
	fmt.Println(fibonacci(10))
}
