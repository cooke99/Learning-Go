package main

import "fmt"

func plus(a int, b int) int {
	// could also do plus(a, b int) as both a and b are ints
	return a + b // Requires explicit returns, won't return last expression by default
}

func plusPlus(a int, b ...int) int {
	for i := 0; i < len(b); i++ {
		a += b[i]
	}
	return a
}

func vals() (int, int) {
	return 3, 7
}

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plusPlus(1, 2, 3, 4, 5))

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

}
