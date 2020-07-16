// syntax: func Sprint(a ...interface{}) string.package main

// golang program to illustrate the usage of fmt.Sprint() function.

package main
// importing fmt, io, and os
import (
	"fmt"
	"io"
	"os"
)

// calling main.
func main()  {
	// declaring some const variables.
	const name, dept = "Julia Veronica", "backend"
	// calling Sprint() function.
	sprint := fmt.Sprint(name, " is a ", dept, " Portal.\n")
	// calling WriteString() function to write the contents of the string "sprint" to "os.Stdout".
	io.WriteString(os.Stdout, sprint)
}
// output: Julia Veronica is a backend Portal.