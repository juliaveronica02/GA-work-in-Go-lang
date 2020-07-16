// _(underscore) in golang is known as blank identifier.
// golang program to show the compiler
// throws and error if a variable is declared but not used.
package main

import "fmt"

func main()  {
	// calling the function.
	// function returns two values which are assigned to mul and div identifier.
	mul, div := mul_div(105,7)
	// only using the mul variable.
	// compile will give an error.
	fmt.Println("105 x 7 = ", mul)
}
// function returning two.
// values of integer type.
func mul_div(n1 int, n2 int) (int,int) {
	// returning the values.
	return n1 * n2, n1 / n2
}
// output: ./main1.go:11:7: div declared and not used.