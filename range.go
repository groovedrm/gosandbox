package main

import "fmt"

func main(){

	// Explicitly setting a slice
	nums := []int{2, 3, 4}
	sum := 0

	// Summing up values in range
	for _, num := range nums{
		sum += num
	}
	// Note: putting the _ at beginning because we don't care
	// about the index in this example, just the actual values
	// of the slice

	fmt.Println("Sum of Range:", sum)

	// Print out values if modulo 2 works
	nums10 := []int{1,2,3,4,5,6,7,8,9,10}
	for idx, val := range nums10 {
		if idx % 2 == 0 {
			fmt.Println("Modulo Val:", val)
			// Note: could also make this the index and not val
		}
	}

	// Iterate on key/value pairs (i.e. maps) and print
	kvs := map[string]string{"a": "apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}
	// Note that if you exclude v and do something like:
	// for k := range kvs {}
	// You can iterate through they keys and write them
	for k2 := range kvs{
		fmt.Printf("%s --> %s\n", k2, kvs[k2])
		// dummy example accessing the values
		// but can just print out the keys.
		// as well
	}

	// Can also do unicode code points (on strings)
	// Not going to think too much about this right now
	// but can refer to 
	// gobyexample.com/range.html
}