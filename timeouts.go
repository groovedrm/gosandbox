package main

import "fmt"
import "time"

func main() {
	
	// First, creating a channel [buffered]
	c1 := make(chan string, 1)

	// Annymous function
	// Within the function, there is a 2 second sleep initiation
	// And then the value is sent to the channel
	// Example is we're doing an externnal call that returns its result
	// on channel c1 after 2 seconds
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Result1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("Time Out 1")
	}

	// In select above, the select is waiting for the first receiver to respond
	// Since the receiver from c1 channel will get a value after 2 seconds, whereas
	// the timeout fires after 1 second, the first received value is the time. So,
	// only the "Time Out 1" value is display from this case

	c2 := make(chan string, 1) // buffered
	
	go func () {
		time.Sleep(time.Second * 2)
		c2 <- "Result2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("Time Out 2")
	}

	// In contrast to above, result 2 does display
	// because the time out case of the select will only
	// fire after 3 sceonds, and it the select receives
	// a value from c2 after 2 seconds
}