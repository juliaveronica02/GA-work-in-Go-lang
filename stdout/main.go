// to add html file.
package main

import "fmt"

func main()  {
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
// run: go run main.go.
// if we want to add html file we can use bellow code.
// went we run go we just add: go run main.go > index.html.
// and check your folder it will auto have index.html file on your folder or project.
// to start html, press go live or ctrl+shift+p and choose live server.