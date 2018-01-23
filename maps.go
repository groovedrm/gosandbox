package main

import "fmt"

func main() {
	
	// Typical key/value structure like in other languages
	// such as dict in Python
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	// Get value for a key
	
	// Note this syntax below doesn't work because it's not 
	// an explicit setting of variable val
	// val = m["k1"]

	// By contrast, this does work
	val := m["k1"]
	fmt.Println(val)

	// Return number of key/value pairs
	fmt.Println(len(m))

	// Learn: Attempt to copy map doesn't work
	// var m2 = make(map[string]int)
	// copy(m2, m)
	// Line 29 Error: arguments to copy must be slices

	// Delete key/value pair from a map
	m3 := make(map[string]int)
	m3["k1"] = 7
	m3["k2"] = 13
	fmt.Println(m3)
	delete(m3, "k2")
	fmt.Println(m3)

	// Optional second return value
	// Explicit use
	err, prs := m["k2"]
	fmt.Println(prs)
	fmt.Println(err)
	// In this example, err returns true [i.e. no error]


	fmt.Println("Ignoring return value below")
	_, prs2 := m["k2"]
	fmt.Println(prs2)
	// Simply ignoring the optional return value

	_, prs3 := m["k3"]
	fmt.Println(prs3)
	// returns false because k3 doesn't exist in map

	fmt.Println("Explicit Use of Both Values Below")
	vl, prs4 := m["k3"]
	fmt.Println(vl)
	fmt.Println(prs4)
	// Returning the vl variable is 0 -- there is no value
	// but since the value of the key is type int, default
	// zero is returned, while prs4 returns false 
	// as the key doesn't exist.

	// Declare and initialize a map in one shot
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)



}