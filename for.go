package main

import "fmt"

func main() {
	// for is Go's only looping construct

	i := 1
	for i <= 3 {
		fmt.Println("i", i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println("j", j)
	}

	for k := range 3 {
		fmt.Println("range", k)
	}

	for {
		// without condition it will loop repeatedly
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n", n)
	}

}
