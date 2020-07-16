// golang program to illustrate the usage of fmt.Sprint() function.package main

// including the main package.
package main
// importing fmt, io and os.
import (
	"fmt"
	"io"
	"os"
)

// calling main.
func main()  {
	// declaring some const variables.
	const num1, num2, num3 = 5, 10, 15
	// calling Sprint() function.
	sprint := fmt.Sprint(num1, num2, num3, "\n")
	// calling WriteString() function to write the contents of the string "sprint" to "os.Stdout".
	io.WriteString(os.Stdout, sprint)
	// output: 5 10 15.

	// 3. example
	// Declaring some const variables 
    const str1, str2, str3 = "a", "b", "c"
  
    // Calling Sprint() function 
    s := fmt.Sprint(str1, str2, str3) 
  
    // Calling WriteString() function to write the 
    // contents of the string "s" to "os.Stdout" 
	io.WriteString(os.Stdout, s) 
	// output: abc.
}
// \n for line break or new line.