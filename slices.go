package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// An unininitialized slice equals to nil and has length 0

	x := make([]string, 3)                                   // create an empty slice with non-zero length.
	fmt.Println("empty:", x, "len:", len(x), "cap:", cap(x)) // cap(x) gives slice capacity, which is by default = to its length
	// if we know the slice is going to grow ahead of time, we can pass a capacity explicitly as an additional parameter to make.

	x[0] = "a"
	x[1] = "b"
	x[2] = "c"
	fmt.Println("x:", x)
	fmt.Println("get:", x[2])
	fmt.Println("len:", len(x))

	// append returns a slice containing one or more new values
	x = append(x, "d")
	x = append(x, "e")
	x = append(x, "f", "g")
	fmt.Println("x:", x)

	y := make([]string, len(x))
	copy(y, x) // Returns number of elements copied - will only copy min(len(y), len(x)) number of elements from x to y
	fmt.Println("y:", y)

	// slice operator supported:
	z := x[2:5]
	fmt.Println("z:", z)
	z = x[:5]
	fmt.Println("z:", z)
	z = x[2:]
	fmt.Println("z:", z)

	t := []string{"g", "h", "i"} // declare and initialize slice in 1 line
	fmt.Println("t:", t)

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // slices can be composed into multi-dimensional data structures - the length of the inner slices can vary unlike multi-dim arrays.
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	underlyingArray := [3]string{"a", "b", "c"}
	slice0 := underlyingArray[0:2]
	slice1 := slice0[:]
	fmt.Println("arr:", underlyingArray, "slice[0:2]:", slice0, "slice1:", slice1)
	underlyingArray[1] = "z"
	fmt.Println("arr:", underlyingArray, "slice[0:2]:", slice0, "slice1:", slice1)
}
