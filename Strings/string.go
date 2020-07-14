package main

import "fmt"

func main ()  {
	// len = length.
	fmt.Println(len("Hello World"))
	// result: 11.

	// individual character in the string.
	fmt.Println("Hello World" [1])
	// result: 101.

	fmt.Println("Hello " + "World")
	// result: Hello World.
}
// run: go run string.go.