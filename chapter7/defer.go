package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second() // defer basically sends second() to the end of main
	first()
}

/*
Usually used for freeing resources after being used. For instance, closing a file after it has been opened.
f, _ := os.Open(filename)
defer f.Close()

Usefull to:
1) keep close next to open
2) if we have multiple returns, close will happen before any of them
3) deferred fuctions run even if a run-time panic error occurs.
*/
