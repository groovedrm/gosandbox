package main

import "fmt"


func f(from string) {
	// When using for loops, don't forget about the semicolon (;) in between
	// the loop paramters.
	for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

	// Here is the synchronous function call
	f("direct")

	// Here is the asynchronous function call
	// f("goroutine")
	go f("goroutine")
	// So, only change is to insert word "go" 
	// just before the function call, which opens up
	// a new "lightweight" thread.

	// The piece below was not explained well.
	// Short story, the goroutine thread runs independent of the main
	// program; however, it will end execution when the main program
	// ends execution. So, if we don't have the block below, the program
	// ends when the f("direct") call completes, even though the 
	// other asynchronous call is executing.

	// As a result, we need something (like below) to make the program
	// "wait" for the asychronous call to finish. Oddly seems like a blocking
	// problem but it's not because you have numerous asynchronous routines
	// See file: goroutines_ex2.go
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}