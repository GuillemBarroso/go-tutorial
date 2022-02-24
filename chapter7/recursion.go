package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // a func that returns itself, thus, recursive
}

func main() {
	fmt.Println(factorial(3))
}
