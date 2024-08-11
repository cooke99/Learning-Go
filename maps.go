package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // make(map[key-type]val-type)
	m["k1"] = 7
	m["k2"] = 12

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) // If key doesn't exist zero-value of the value type is returned

	fmt.Println("len(m):", len(m)) // # key-val pairs

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m) // Deletes all key-value pairs
	fmt.Println("map:", m)

	_, present := m["x"] // optional 2nd return val when getting a value from a dict indicates if key is present in the map
	fmt.Println("present:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)

	var n2 = map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
