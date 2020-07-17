// buffalo scanners and how to use scanners.package main
// because we will use that to read from the incoming connection.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)
func main()  {
	// we are going to listen and then right we defer to close of that listener,
	// we get a listener and then we are eternally looping over.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		// accept and got a connection.
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// we going launch a go routine.
		go handle(conn)
	}
}
// after lauch we pass this off to below function with the name handle.
// at here we goint to handle this connection and also passed our connection into that.
func handle(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	// we never get here.
	// we have an open stream connection.
	// how foes the above reader know when its's done?.
	fmt.Println("Code got here.")
}