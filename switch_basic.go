package main

import "fmt"
import "time"

func main() {

	switch time.Now().Weekday() {
		case time.Saturday:
			fmt.Println("Saturday")
		case time.Sunday:
			fmt.Println("Sunday")
		default:
			fmt.Println("Not a weekend")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			// Note: the .type notation can only be used within or in conjunction with a switch statement
			case bool:
				fmt.Println("Boolean variable")
			case int:
				fmt.Println("Integer variable")
			case float64:
				fmt.Println("Floating-point variable")
			default:
				fmt.Println("Cannot determine variable type %T\n", t)
		}

	}

	whatAmI("Hi")
	whatAmI(true)
	whatAmI(false)
	whatAmI(3)
	whatAmI(3.50)
}