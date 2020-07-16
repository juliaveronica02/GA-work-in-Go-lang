package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
func main()  {
	name := "Julia Veronica"
	str := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World </title>
		</head>
		<body>
		<h1> ` + name + ` </h1>
		</body>
		</html>
		`)
		newFile, err := os.Create("index.html")
		if err != nil {
			log.Fatal("error creating file", err)
		}
		defer newFile.Close()
		io.Copy(newFile, strings.NewReader(str))
}
// run: go run main.go.
// after run we can see we already have index.html file.
// to make sure we can delete the existing index.html file to see the code success or not.

// note:
// sprintf and printf functions work in similiar fashion.
// sprint does everything what print does but does not write the resulting string to standard output,
// instead, it returns it.
// simple one for spirnt is formats the string and does not print to console. It returns the formatted string.
// Printf formats the string and prints in the console. It does not return the formatted string.