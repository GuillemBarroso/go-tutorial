package main

import "fmt"

var numberName string

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			numberName = "Fizz"
		}
		if i%5 == 0 {
			numberName += "Buzz"
		}
		fmt.Println(i, numberName)
		numberName = ""
	}
}
