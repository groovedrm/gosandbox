package main

import "fmt"

func main() {
	
	// --------------------------------------------------------------------------------
	// Arrays
	// --------------------------------------------------------------------------------

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

	// --------------------------------------------------------------------------------
	// Slices [powerful interface to sequences than arrays in Go]
	// --------------------------------------------------------------------------------

	// Want to better understand syntax of slice vs. array
	// Perhaps syntatically different so you can actually visually see difference
	// between the two
	slc := make([]string, 3)
	fmt.Println("Slice: ", slc)
	
	slc[0] = "hi"
	slc[2] = "Chris"
	fmt.Println("Slice Revised: ", slc)

	// Can use the append statement to add elements to slice
	slc = append(slc, ", Nice to meet you")
	fmt.Println("Slice Revised 2: ", slc)

	slc = append(slc, ". Welcome!")
	fmt.Println("Slice Revised 3: ", slc)

	// Copying a slice accomplished below.

	slc_c := make([]string, len(slc))
	// Using the len of the slc slice to bound this one.
	// Also note the use of len works with slices 
	// much like it works with arrays.
	
	copy(slc_c, slc)
	// Copy syntax (copy to, copy from)
	fmt.Println("Copy: ", slc_c)

	// Can access a portion of a slice
	fmt.Println(slc[2:3])
	fmt.Println(slc[2:])
	fmt.Println(slc[:3])

	// Reminder: initialize array in one line.
	initarray := [5]int{1,2,3,4,5}
	initslice := []int{1,2,3,4,5}

	fmt.Println(initarray)
	fmt.Println(initslice)
	// Takeaway: slice doesn't need the length
	// in the [] space when initializing


	// Also, can initialize 
	twoDD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoDD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoDD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoDD)
    // Takeaway, the length of inner
    // slice can be different than other slices
    // Output: [[0] [1 2] [2 3 4]]


}
 