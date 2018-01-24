package main

import "fmt"
import "time"

// We'll run this function in a goroutine
// The channel "done" will be used to notify
// another goroutine that this functions
// work is done.
// In this case, the goroutine we are notifying
// and blocking is going to be the main() function
// i.e. main goroutine.
func worker(done chan bool) {
	fmt.Print("working..")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1) 
	// Note, buffered
	// Meaning we don't specifically need a receiver

	go worker(done)

	// The code below blocks main from completion 
	// until a notification is received from the channel
	// via the line done <- true above.
	<-done

}