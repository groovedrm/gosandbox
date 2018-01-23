package main

import "fmt"

type rect struct {
	width, height int
}

// This is an example of a pointer receiver
// because of the * syntax next to rect
func (r *rect) area() int{
	return r.width * r.height
}

// This is an example of a value receiver
// as there is no * and instead a normal variable
// parameter is passed

func (r rect) perim() int{
	return 2*r.width + 2*r.height
}


func main() {

	shape := rect{width: 10, height: 5}

	fmt.Println("Area:", shape.area())
	fmt.Println("Perim:", shape.perim())

	rp := &shape
	fmt.Println("Pointer Area: ", rp.area())
	fmt.Println("Pointer Perim: ", rp.perim())
	
}

// Other Notes
// Using a pointer receiver type allows you to mutate the 
// receiving struct, otherwise you are making a copy of the
// struct. Not a problem if you want to return it and use
// it in another variable, but will have to actually make that assignment
// vs the pointer that can mutate the struct directly.