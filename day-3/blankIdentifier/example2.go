// golang program to the use of blank identifier.package blankIdentifier

package main

import "fmt"

// main function.
func main() { 
	// calling the function.
	// function returns two values which are.
	// assigned to mul and blank variable.
	mul, _ := mul_div(105,7)
	// only using the mul variable.
	fmt.Println("105 x 7 =", mul)
}
// function returning two.
// values of integer type.
func mul_div(n1 int, n2 int) (int, int) {
		// returning the values.
		return n1 * n2, n1 / n2
}
// output: 105 x 7 = 735.

// note:
// we can use multiple blank identifier in the same program. 
// we can say a golang program can have multiple variable using the same identifier name which is the blank identifier.
// we can use any value of any type with the blank identifier.