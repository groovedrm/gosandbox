package main

import "fmt"

// Forming closures in Go
// Allows use of anonymoous functions

// function intSeq returns another function
// defined anonymously in intSeq in "return func" line
// A closure means that the anonymous functions

func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
			return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	
}