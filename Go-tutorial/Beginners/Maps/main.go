/*
// maps:
	-> Data-structure that holds key-value pair:


*/

package main

import "fmt"

func main() {

	// Declare with make:
	m1 := make(map[string]int)
	m1["Sahil"] = 1
	m1["Prince"] = 2
	fmt.Println(m1)

	// Map literals:
	m2 := map[string]string{
		"lang": "Go",
		"os":   "mac",
	}
	fmt.Println(m2)

	// Accessing & Updating:
	m := map[string]int{"A": 1, "B": 2}
	fmt.Println(m["A"]) // assessing "A"

	m["B"] = 99 // Updating
	fmt.Println(m)

	// Accessing non-existing key:
	fmt.Println(m["C"]) // 0 default value

	// Check if key exist:
	val, ok := m["C"]
	if ok == true {
		fmt.Println("Value exist: ", val)
	} else {
		fmt.Println("Key Not Exist")
	}

	// Delete key:
	delete(m, "B")
	fmt.Println(m)

	// Iterating in map:
	for key, val := range m {
		fmt.Println(key, ":", val)
	}

	// Insert key:
	m["B"] = 88
	m["C"] = 99
	fmt.Println(m)

	// map of slice:
	friends := map[string][]string{
		"Prince": {"Radharaman", "krishna", "Radha", "Gopinath", "Bihariji", "Radhavallav"},
		"Sahil":  {"jai-jai", "Thakurji", "madhav", "lala"},
	}

	fmt.Println(friends)
	fmt.Println(friends["Prince"])

	// nil map vs empty map:
	var maap1 map[string]int      // nil map, cannot write
	maap2 := make(map[string]int) // empty map, can write

	fmt.Println(maap1 == nil)
	fmt.Println(maap2 == nil)

}
