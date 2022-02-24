package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // keeps generating even numbers
	fmt.Println(nextEven()) // keeping the previous value of i
	fmt.Println(nextEven()) // since it is a func that returns a func
}
