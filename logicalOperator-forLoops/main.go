package main

import "fmt"

func main()  {
	// logical operators.
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true= ", !true)

	// for loops.
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// shorthand for i = i +1.
		i++
	}
	// relational operators include ==, !=, <, >, <=, and >=.
	// we can also define a for loop like code bellow, but we need semicolons.
	for j := 0; j <5; j++ {
		fmt.Println("j",j)
	}
}