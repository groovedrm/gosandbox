package main

import "fmt"

func main() {
	
	// Make two channels, not buffered
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages: // receiver
		fmt.Println("Rec. Message: ", msg)
	default:
		fmt.Println("No message received")
	}
	// This is a non-blocking receive. If there is a value 
	// available on messages (and received in the select), it will
	// be printed, otherwise default case is printed

	// Similarly, a non-blocking send
	msg := "hi"
	select {
	case messages <- msg: // sender
		fmt.Println("Sent Message: ", msg)
	default:
		fmt.Println("No Message Sent")
	}

	// 
	select {
	case msg := <-messages:
		fmt.Println("Received Message: ", msg)
	case sig := <-signals:
		fmt.Println("Received Signal: ", sig)
	default:
		fmt.Println("No activity")
	}
}