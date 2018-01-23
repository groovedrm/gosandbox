package main

import "fmt"

// Variadic functions accept an unlimited number of arguments
// In their parameter set
// Using ... syntax

func totals(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}


func main() {
	// Declare a couple numbers
	n1 := 3
	n2 := 5
	sum := totals(n1, n2)
	fmt.Println(sum)

	n3 := 142
	sum2 := totals(n1, n2, n3)
	fmt.Println(sum2)
	// Need to learn more about the no new variables
	// error to left of := sign
	// Which happens when you use sum variable in both 
	// lines 23 and 27

	// Example with a slice of varying length
	ns := []int{150, 150, 150}
	fmt.Println(totals(ns...))
	// Note: use the ... after the varying length
	// variable to sum over the contents of the variable
	
}