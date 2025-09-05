/*
// Array:
	-> Fixed size sequence of same type
	-> It's same same as other language

*/

package main

import "fmt"

func main() {

	// Declare array with size:
	var arr1 [3]int
	fmt.Println(arr1)

	// Declare array with values:
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Let Go count the size, Compiler will count the size from the value
	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	// Declare Array & take input from user:
	var arr4 [5]int
	for i := 0; i < 5; i++ {
		arr4[i] = i * i
	}

	// Print array value:
	for _, val := range arr4 {
		fmt.Print(val, " ")
	}
	fmt.Println()

	// Multi-Dimensional Array:
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
}
