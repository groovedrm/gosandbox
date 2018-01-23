package main

import "fmt"

// Int parameter
// So arguments passed by value
// the ival in the function is distinct
// from the calling function
func zeroval(ival int) {
	ival = 0
}

// Has the * syntax which means accepts 
// an int pointer rather than an actual int value
// Then *iptr deferences the pointer from its memory address
// to the current value at that address
// Assigning a value to the deferenced pointer changes
// the value associated with the memory address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("Initial: ", i)

	zeroval(i)
	fmt.Println("ZeroVal:", i)
	// Note, the variable in function is distinct from
	// variable value passed, so nothing changes with I

	// the & syntax is how we reference the memory
	// address of the pointer
	zeroptr(&i)
	fmt.Println("ZeroPtr:", i) 
	fmt.Println("Pointer: ", &i)
	
}