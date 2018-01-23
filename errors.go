package main

import "errors"
import "fmt"

// The convention for errors is they are propagated through the functionc all as the last
// value returned, and have type error, a standard interface in the Go language
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
		// errors.New connstructs a basic error with a message 
		// returned back to the console (in this case)
	}

	return arg + 3, nil
	// Here we are returning two values as well; however, the nil in place
	// of an actual error is used because there is no actual error value returned
}

// It's possible to use custom error types
// as errors (the standard type) by implementing the error() 
// method on them. This is an example that explicitly uses a custom
// type
type argError struct{
	arg int
	prob string
}

// Custom function implementing error
// Note that we are using a pointer argument instead of a value
// argument to explicit change the value of the struct, rather
// than making a copy of the struct
func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with value"}
		// Notes: above, we are using the a pointer to the
		// error struct in order to make sure we deliberate mutate
		// that struct rather than make a copy, and also assigning 
		// a value prob variable of the struct (within the returrn)
	}

	return arg + 3, nil
}

func main() {

	// Note: using the _ below because we don't
	// care about the index of the array, just the array
	// value on its own

	for _, i := range []int{7, 42} {

        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // Notes: 
    // the if r, e := f1(i); e != nil 
    // error check on the same line syntax is a
    // common idiom in Go code

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }

    // If you want to programmatically use the data in a custom error, 
    // you’ll need to get the error as an instance of the custom 
    // error type via type assertion
}