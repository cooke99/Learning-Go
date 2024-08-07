package main

import "fmt"

func main() {
	var a = "initial" // var declares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // can declare multiple variables at once
	fmt.Println(b, c)

	var x, y = 1, "a" // can mix types
	fmt.Println(x, y)

	/*
		Throws syntax error:
		var p, q int = 1, "a"
		fmt.Println(p, q)
	*/

	var d = true
	fmt.Println(d)

	var e int // uninitialized variables are 0-valued
	fmt.Println(e)

	f := "apple" // Shorthand for var f string = "apple" --> only available inside functions
	fmt.Println(f)

}
