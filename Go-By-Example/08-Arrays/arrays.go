package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("array a: ", a)

	a[4] = 100 // we can set value at in index

	fmt.Println("array a: ", a)
	fmt.Println("default value at index 3: ", a[4])

	fmt.Println("length of an array a: ", len(a)) // we can check the length of an array using len() function.

	b := [5]int{1, 2, 3, 4, 5} // Syntax to initialize an array in single line
	fmt.Println("array b: ", b)

	// If we don't want to initialize the size by itself, want to count by the compiler then:
	c := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array c: ", c)
	fmt.Println("length of an array c: ", len(c))

	// If we specify the index with :, then index b/w them will be '0'
	d := [...]int{100, 3: 400, 500, 500}
	fmt.Println("array d: ", d)

	// Array are one dimensional, but you can compose type to build multi-dimensional data-structure.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	// Print 2D array:
	fmt.Println("2D array: ", twoD)
	for i := range 2 {
		for j := range 3 {
			fmt.Print(twoD[i][j], " ")
		}
		fmt.Println()
	}

	// We can initialize the previous declared multi-dimensional array at once:
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("2D array second : ", twoD)

	// Declare & initialize at once:
	twoD2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("2D array: ", twoD2)

}
