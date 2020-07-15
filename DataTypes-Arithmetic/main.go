// Every go program starts with a package declaration which provides a way for use to reuse code.
package main

// import allows use to use code from another package.
// the format package provides formatting for input and output.
import "fmt"

// a comment for one line.

/*
multiline comment.

*/

// function start with func and surround the code with {}.
// main is the function that is executed when we executed our program.
func main()  {
	// Println is a function in the fmt package that outputs a string, which -
	// is surrounded by double quotes and a newline to the screen.
	fmt.Println("Hello World")

	// we can get a description of a function by typing godoc fmt Println -
	// for example in the terminal.

	// execute go programs in the terminal just write go run fileName.go.

	// variables are statically typed, which means their type can't change.
	// variable names must start with a leeter and mau contain letters, numbers or the strip (-)/ underscore (_).

	// an int is a positive or negative number without decimals.
	// Versions:
	// uint8: unsigned 8bit integers (0 to 255).
	// uint16: unsigned 16bit integers (0 to 65535).
	// uint32: unsigned 32bit integers (0 to 4294967295).
	// uint64: unsigned 64bit integers (0 to 18446744073709551615).
	// int8: signed 8bit integers (-128 to 127).
	// int16: signed 16bit integers (-32768 to 32767).
	// int32: signed 32bit integers (-2147483648 to 2147483647).
	// int64: signed 64bit integers (-9223372036854775808 to 9223372036854775807).

	var age int64 = 40
	
	// a float is a number with decimals.
	// versions: float32, float64.
	var favNum float 
}