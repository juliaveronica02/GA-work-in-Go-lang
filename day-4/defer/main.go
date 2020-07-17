package main

import (
	"fmt"
	"io"
	"log"
	"os"
)
func main()  {
	// 1. example.
	// defer statement is executed, and places.
	// fmt.Println("Bye") on a list to be executed prior to the function returning.
	defer fmt.Println("Bye")
	// the next line is executed immediately.
	fmt.Println("Hi")
	// fmt.Println*("Bye") is now invoked, as we are at the end of the scope function.
	// output: Hi and Bye.

	// 2. example.
	// defer to cleanup resources. code below to create readme.txt file.
	if err := write("readme.txt", "this is readme file"); err != nil {
		log.Fatal("failed to write file:", err)
	}
}
// function write for 2. example.
func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// defer file.Close() to tells the compiler that it should execute the file.Close prior to exiting the function write.
	 defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return nil
	// we also can use: return file.Close() or defer file.Close().
}
// run: go run main.go.
// result of defer cleanup resources: have readme.txt after run.