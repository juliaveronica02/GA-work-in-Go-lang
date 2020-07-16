package main

import "fmt"

func main(){
	listOfNums := []float64{1,2,3,4,5}
	fmt.Println("Sum :", addThemUp(listOfNums))
	// get 2 values from a function.
	num1, num2 := next2Values(5)
	fmt.Println("number:",num1, num2)
	// send an undefined number of values to a function (variadic funtion).
	fmt.Println(subtractThem(1,2,3,4,5))
	// we can create a function in a function.
	// it has access to the local variables of the containing function.
	// a function like this with no local varibles is a closure.
	num3 := 3
	doubleNum := func () int  {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	// calling a recursive funtion.
	fmt.Println(factorial(3))
	// defer executes a function after the inclosing funtion finishes.
	// defer can be used to keep functions together in a logical way.
	// but at the same time execute one last as a clean up operation.
	// example. defer closing a file after we open it and perform operations.
	defer printTwo()
	printOne()
	// use recover() to catch a division by 0 error.
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))
	// we can catch our own errors and recover with panix and recover.
	demPanic()
}

// functions allow us to reuse code and provide structure.
// func funcName(parametersPassed) return type.
// functions don't have access to any variables aside from those and passed into it.
func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		// shorthand for sum = sum +val.
		sum += val
	}
	return sum
}

// go function can return multiple values.
func next2Values(number int) (int,int) {
	return number+1, number+2
}

// we can receive an undefined number of values with args ...int.
func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args{
		finalValue -= value
	}
	return finalValue
}

// example of recursion : function calls itself.
// factorial(3)
// 3 * factorial(2) == 3 * 2.
// 2 * factorial(1) == 2 * 1.
// factorial(0) == 1.
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num -1)
}

// used to demonstrate defer.
func printOne()  {
	fmt.Println(1)
}
func printTwo()  {
	fmt.Println(2)
}

// if an error occurs we can catch the error with recover and allow code to continue to execute.
func safeDiv(num1, num2 int) int {
	defer func ()  {
		fmt.Println(recover())
	}()
	solution := num1/num2
	return solution
}

// demonstrate how to call panic and handle it with recover.
func demPanic()  {
	defer func ()  {
		// if we didn't print the message nothing would show.
		fmt.Println(recover())
	}()
	panic("Panic")
}