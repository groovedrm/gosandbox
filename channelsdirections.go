package main

import "fmt"


// When using channels as function parameters, you can specify if a 
// channel is meant to only send or receive values. This specificity 
// increases the type-safety of the program.

// Function only accepts a channel for sending values
// Syntax for accepting sender is chan to left of <-
// but datatype to the right
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// By contrast, this fuction accepts two channels
// The first is a receiver via the chan to right
// of the <-, and the second is a sender with the 
// chan to the left of the <-. 
// In both cases, data type still to the right.
func pong(pings <- chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	// First, make the channels
	// Note we don't pass values in this first piece
	// Just the channel parameters
	// Both buffered so we don't have to worry about
	// receivers just yet
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Now the functions
	ping(pings, "passed message")
	// Calling ping function, passing
	// in the channel as well as the message

	pong(pings, pongs)
	// First argument receives channel value
	// from pings, whereas the second sends the 
	// channel value

	// Finally, receiver for value from pongs
	fmt.Println(<-pongs)

	// Need to learn more about this:
	// https://gobyexample.com/channel-directions

}
