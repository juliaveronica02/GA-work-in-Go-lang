package main

import "fmt"

func main()  {
	// an array holds a fixed number of values of the same type.

	var favNums[5] float64
	favNums[0] = 163
	favNums[1] = 78557
	favNums[2] = 691
	favNums[3] = 3.141
	favNums[4] = 1.618

	// we can access the value by supplying the index number.
	fmt.Println(favNums[3])

	// another way of initializing an array.
	favNums2 := [5]float64 {1,2,3,4,5}
	// how to iterate through an array (use _ if a value is not used).
	for i, value := range favNums2 {
		fmt.Println(value, i)
	}
	// slices are like arrays but we leave out the size.
	numSlice := [] int {5,4,3,2,1}
	// we can create slice by defining the first index value to take through the last.
	fmt.Println("numSlice[0]", numSlice[0])
	// if we don't supply the first index it defaults to 0.
	// if we don't supply the last index it defaults to max.
	fmt.Println("numSlice[:2] =", numSlice[:2])
	fmt.Println("numSlice[2:] =", numSlice[2:])
	// we also create an empty slice and define the data type,
	// length (receive value of 0), capacity (max size).
	numSlice2 := make([]int, 5, 10)
	// we can copy slice to another.
	copy(numSlice2, numSlice)
	fmt.Println(numSlice2[0])
	// append values to the end of a slice.
	numSlice2 = append(numSlice2, 0, -1)
	fmt.Println(numSlice2[6])
}
// run: go run main.go.