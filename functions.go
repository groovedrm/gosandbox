package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// Notes:
// 1: Go requires explicit returns
// It won't automatically return the value of the last expression

func plus2(a, b, c int) int {
	return a + b + c
}

// 2: Multiple parameters of same type 
// can be shorthand expressed like the parameters in 
// function plus2 vs. plus

// 3: Go functions support multiple return values
// You simply use a () in the function return type
func plus3(a, b int) (int, float64) {
	return a + b, float64(a + b) / float64(2)
	// Above, working.
	// Sort of like SQL where you need to be deliberate
	// about variable type in mathematical operations
}

func main() {
	
	// First print out straight up function call
	fmt.Println(plus(3,5))
	fmt.Println(plus2(3,5,2))

	ret := plus(3,5)
	fmt.Println("Returned Value: ", ret)

	add, div := plus3(3,4)
	fmt.Println("Return Values: ", add, div)

	_, c := plus3(3,4)
	fmt.Println("Just second parameter: ", c)
}