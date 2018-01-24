package main

import "fmt"

func main() {
	
	// Make a channel, but can buffer up to two values
	messages := make(chan string, 2)
	// An unbuffered default channel looks like this:
	// messages := make(chan string)
	// So we can see the syntax difference
	// Channels are unbuffered by default so can only 
	// acept send values if there is a corresponding
	// receiver.

	// Buffered channels accept a limited number of values without
	// corresponding receivers for those values

	messages <- "buffered"
	messages <- "channel"
	// In first pass of this, we don't need any receivers below
	// can do "go run channelsbuff.go" and no errors are returned

	// But instead if we want to print out the messages
	// sent to the channel, we can add receivers below.

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}