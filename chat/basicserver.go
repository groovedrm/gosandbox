package main

import "fmt"
import "net"
import "bufio"


// func handleConnection(c conn ) {
// 	fmt.Println("Accepting connection", c)
// }

type client struct {
	clientid net.Conn
	clientAddr net.Addr

}

func clientIntro(c net.Conn) {
	w := bufio.NewWriter(c)

	_, err := w.Write([]byte("Welcome!"))
	if err != nil {
		fmt.Println("Msg Delivered")
	} else {
		fmt.Println("Delivery error:", err)
	}

}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Server initialization error")
	}

	// Make a slice of client structs
	// clients := make([] client, 1)
	// clients := []client{}
	clients := []*client{}
	// Length 1 for now, then expand

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
		} else {

			// clientConn := client{clientid: conn, clientAddr: conn.LocalAddr()}
			clientConn := new(client)
			clientConn.clientid = conn
			clientConn.clientAddr = conn.LocalAddr()
			go clientIntro(conn)

			clients = append(clients, clientConn)

			// fmt.Println("Connection: ", conn)
			// fmt.Println(": Local Address: ", conn.LocalAddr())
			// Notes: cannot slice conn
			fmt.Println("New connection")
			fmt.Println("-- Connected clients: ", len(clients))
			fmt.Println("-- Most Recent Client: ", clients[len(clients) - 1: len(clients)])
		}
	}

}