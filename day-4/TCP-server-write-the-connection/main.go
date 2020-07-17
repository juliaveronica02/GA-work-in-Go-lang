package main

import (
	"fmt"
	"io"
	"log"
	"net"
)
func main()  {
	// 8080 is port.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		// Conn: read and write text.
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// WriteSting ask for a writer and the string right soil or write a string to a writer.
		io.WriteString(conn, "\n Hello from TCP Server \n")
		// we can pass that connection here because as print line once a writer after enough once a writer writes.
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")
		// we could write all that stuff back to our connection to close our connection and go back to the listening.
		conn.Close()
	}
}
// before run we must install telnet.
// so, open linux terminal and type: sudo npm -g i telnet.
// sudo npm -g i telnet is means to install telnet in globally.
// and back to visual studio code and run: go run main.go.
// after that klik + on vs code terminal to add new terminal.
// go to file folder and just type: telnet localhost 8080.
// output: 
// Hello from TCP Server .
// How is your day?.
// Well, I hope!Connection closed by foreign host.
// if we close go run, and back to telnet the result is: we can't find the connection.