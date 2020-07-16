package main

import (
	"log"
	"os"
	"text/template"
)
func main()  {
	// show output on terminal when transfer file from tpl.gogtml.
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}