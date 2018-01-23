package main

import "fmt"

func main() {
	
	// Initializes an array with 5 elements -- all zeroes
	var a[5] int
	fmt.Println("Array: ", a)

	// Initializes an array of 5 empty strings
	var b[5] string
	fmt.Println("Array String: ", b)

	// Initializes an array of 5 bool variables, all false
	var c[5] bool
	fmt.Println("Bool Array: ", c)

	// Initialize an array in one line.
	d := [5]int{1,2,3,4,5}
	fmt.Println("Initialized Int Array: ", d)

	e := [5]string{"hi ","my ","name ","is ","chris"}
	fmt.Println(e)

	var twoD [2][3]int
	fmt.Println(twoD)
	fmt.Println(len(twoD))
}