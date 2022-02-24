package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person // use Person type and don't give it a name
	Model  string
}

func main() {
	a := new(Android)
	a.Person.Talk()
}

// the logic behind embedded methods is:
// people can talk -> android is a person -> android can talk
