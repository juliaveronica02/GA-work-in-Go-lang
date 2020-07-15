// to add html file.
package main

import "fmt"

func main()  {
	// first example.
	// long.
	// var name = 
	// shorthand.
	name := "Julia Veronica"
	// type html on bellow.
	tpl := `
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
	`
	fmt.Println("show",tpl)
}