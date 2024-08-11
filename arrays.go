package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	var c = [5]int{1, 2, 3, 4, 5}
	fmt.Println("c:", c)

	b = [...]int{1, 2, 3, 4, 5} // Compiler counts number of elements
	fmt.Println("b:", b)

	// var d = [5]{1,2,3} --> can't infer type from elements

	d := [...]int{100, 3: 400, 500} // if you specify the index with :, the elements in-between will be zeroed
	fmt.Println("d:", d)            // output is d: [100 0 0 400 500]

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("twoD:", twoD)

}
