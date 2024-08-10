package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// can use commas to separate multiple conditions in case statement:
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday!")
	}

	t := time.Now()
	// Can use switch-case without an expression to behave like if/else:
	switch {
	case t.Hour() < 12:
		// Non-constant expression for case expression
		fmt.Println("It's before noon!")
	default:
		fmt.Println("It's after noon!")
	}

	// Type switch compares types instead of values
	// Can use this to discover the type of an interface value
	// Variable t will have type corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(6)
	whatAmI(6.0)
	whatAmI("hey")
}
