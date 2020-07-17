package main

import "fmt"
func main()  {
	// defer statement is executed, and places.
	// fmt.Println("Bye") on a list to be executed prior to the function returning.
	defer fmt.Println("Bye")
	// the next line is executed immediately.
	fmt.Println("Hi")
	// fmt.Println*("Bye") is now invoked, as we are at the end of the scope function.
}