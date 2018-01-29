package main

import "fmt"
import "net"
import "bufio"

// Struct Outlining The Client
/*
type Client struct {
	clientAddr net.Addr
	in_msg chan string
	out msg chan string
	reader *bufio.Reader
	writer *bufio.Writer
}
*/

type Client struct {
	clientAddr net.Addr
	in_msg chan string
	out_msg chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
}

// Func to create and return a new client
// Pointer to the client, return than a full memory copy of client
// * Deferences the client from it's memory
// Using hte pointer because we don't need to keep a copy of each new client
func NewClient(c net.Conn) *Client {
	clientAddr := c.LocalAddr()
	writer := bufio.NewWriter(c)
	reader := bufio.NewReader(c)

	// Setting the return values
	client := &Client{
		clientAddr: clientAddr,
		in_msg: make(chan string),
		out_msg: make(chan string),
		writer: writer,
		reader: reader,
	}

	return client
}
/*
-- My Old Intro Code
func clientIntro(c net.Conn) {
	w := bufio.NewWriter(c)

	fmt.Println(w)
	wrt, err := w.Write([]byte("Welcome!"))
	if err != nil {
		fmt.Println("Msg Delivered")
	} else {
		fmt.Println(wrt) // seesm to be printing 8 bytes
		fmt.Println("Delivery error:", err)
	}
}
*/

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Server initialization error")
	}

	// Running slice of clients
	clients := []*Client{}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
		} else {
			clientConn := NewClient(conn)
			clients = append(clients, clientConn)

			fmt.Println("New connection")
			fmt.Println("-- Connected clients: ", len(clients))
			fmt.Println("-- Most Recent Client: ", clients[len(clients) - 1: len(clients)])
		}
	}

}