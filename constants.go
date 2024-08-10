package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// x := "a" Constants can not be declared using the := syntax

func main() {
	fmt.Println(s)

	const n = 500000 // const can appear anywhere a var can

	const d = 3e20 / n // Constant expressions perform arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) // A numeric constant has no type until it's given one such as by explicit conversion

	fmt.Println(math.Sin(n)) // A number can be given a type by using it in a context that requires one e.g. variable assignment or function call
	// math.Sin expects float64

	// d = 1

}
