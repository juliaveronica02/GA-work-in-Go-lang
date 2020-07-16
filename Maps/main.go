package main

import "fmt"
func main()  {
	// a map is a collection of key value pairs.
	// created with varName := make(map[keyType] valueType).
	presAge := make(map[string] int)
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println("age 42:",presAge["TheodoreRoosevelt"])
	// get the number of items in the map.
	fmt.Println("presAge:",len(presAge))
	// the size changes when a new item is added.
	presAge["John F. Kennedy"] = 43
	fmt.Println("age 43:",len(presAge))
	// we can delete by passing the key to delete.
	delete(presAge, "John F. Kennedy")
	fmt.Println("delete:",len(presAge))
}
// run: go run main.go.