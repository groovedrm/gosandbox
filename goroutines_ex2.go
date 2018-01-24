package main

import "fmt"


func f(from string) {
	// When using for loops, don't forget about the semicolon (;) in between
	// the loop paramters.
	for i := 0; i < len(from); i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

	// Async calls below where words are different length
	// so we can see them interwoen in output
	go f("superlongwordtotestoutput_blahblahblah")
	go f("notaslongword")
	go f("go")

	go func(msg string) {
        fmt.Println(msg)
    }("interwoven")

    // Depending on the execution, the execution of all this changes. Very interesting.
    // Need to learn more about the execution stack.
    // https://blog.golang.org/go-concurrency-patterns-timing-out-and

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}