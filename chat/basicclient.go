package main

import "fmt"
import "net"
import "bufio"



func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server")
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")


	rc := []byte{}
	msg, err := bufio.NewReader(conn).Read(rc)
	if err != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Connected, but error:", err)
	}
	// status, err := bufio.NewReader(conn).ReadString('\n')
}