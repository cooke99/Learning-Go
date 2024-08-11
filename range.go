package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// range on arrays and slices provides both idx and value for each entry
	for idx, num := range nums {
		fmt.Println("idx:", idx, "val:", num)
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on maps iterates over key, value pairs
	dict := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range dict {
		fmt.Printf("%s -> %s\n", key, val)
	}

	// Can just iterate over keys too:
	for key := range dict {
		fmt.Printf("%s\n", key)
	}

	// range on strings iterates over Unicode code points - first value is the starting byte index of the rune and the second is the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
