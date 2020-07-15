package main

import "fmt"

func main() {
	// if statement.
	myAge := 19
	if (myAge >= 16) {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}

	// we can use else if perform different actions, but once a match -
	// is reached the rest of the conditions are not checked.
	if (myAge >= 16) {
		fmt.Println("You can drive")
	} else if myAge >= 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("You can have fun")
	}

	// switch statements are used when we have limited options.
	switch myAge {
	case 16 : fmt.Println("Go Drive")
	case 18 : fmt.Println("go vote")
	default : fmt.Println("Go have fun")
	}
}