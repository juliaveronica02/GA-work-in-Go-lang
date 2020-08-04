package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	fmt.Println("this is main function")
	// put goroutine on bottom because the file always execute from top.
	// if we put on top it's okay to because it will be last program execute.
	go hello() // call hello function.
}
