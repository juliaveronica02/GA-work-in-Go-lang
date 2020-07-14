package main

import "fmt"

func main()  {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true && false)
	fmt.Println(!true)
}
// result: true, false, true, true, false.
// run: go run bool.go.