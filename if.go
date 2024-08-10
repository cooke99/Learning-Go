package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 7 or 8 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is neagtive")
	} else if num < 10 {
		fmt.Println(num, "is < 10")
	} else {
		fmt.Println(num, "> 10")
	}
	// any variables declared in statement are available in current and subsequent branches i.e.
	// num available in if/else if/else branches but not outside
}
