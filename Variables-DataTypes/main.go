package main

import "fmt"

func main()  {
	// Main types:
	// String.
	// Bool.
	// Int.
	// int8, int16, int 32 and int64.
	// float32 and float 64.

	// 1. example using var (string).
	// the green text it's not error, this is data type (string).
	var myName string = "Julia Veronica"
	fmt.Println("name: ", myName)

	// 2. example using var (string).
	var name = "Juve"
	fmt.Println("short name: ", name)
	// result: short name:  Juve.

	// 3. example using var (string, integer (int)).
	var firstName = "Julia"
	var age int = 19
	fmt.Println("First Name and age: ", firstName, age)
	// result: First Name and age:  Julia 19.

	// 4. example using var (string, integer (int)).
	var firstNames = "Julia"
	var ages = 19
	fmt.Println("First Name and age: ", firstNames, ages)
	// result: First Name and age:  Julia 19.

	// 1. example int.
	fmt.Println("1 + 1 = ", 1+1)
	// result: 1 + 1 =  2.

	// 2. example int.
	fmt.Println("result 1 + 1 = ", 1.0 + 1.0)
	// Notice that we use the .0 to tell Go that this is a floating point number instead of an integer.
	// Running this program will give you the same result as before.
	// result: result 1 + 1 =  2.
	
}
// run: go run main.go.
// result: name:  Julia Veronica.