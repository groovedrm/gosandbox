package main

import "fmt"

func main() {
	// Channels are pipes that connect concurrent 
	// goroutines. You can send values into channels from one
	// goroutines and receive those values into another goroutine


	// Here, we are creating a new channel using
	// make command, and then (chan <data type>)
	messages := make(chan string)

	// You send a value into a channel using 
	// channel <- syntax
	// So in example below, sending "ping" to 
	// messages channel made above using an 
	// anoynmous goroutine

	go func() {messages <- "ping"}()

	// The syntax below receives a value from the channel
	// So we want to receive the ping message below.
	// When we run, ping is passsed from one goroutine to
	// another via the channel

	msg := <-messages
	fmt.Println(msg)

	// By default, sends/receives block until both the
	// sender and receiver are ready. This property allowed us
	// to wait at the end of our program for the ping message
	// without having to use any other sycrhonization
}