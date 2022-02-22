package main

import "fmt"

func main() {
	// definition variable as string and modification
	var x string
	x = "first "
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	// automatic variable type assignation
	y := "Hello World"
	fmt.Println(y)

	// constants
	const NUMBER int = 5
	fmt.Print(NUMBER)
}
