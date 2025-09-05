/*

// Iterations:
	-> Go has only "for"
	-> Syntax:
		-> Initialization: initialize anything necessary in iterations
		-> Condition: check after every iteration
		-> Post: runs at the end of iterations.

		for initialization; condition; post {
			// statements to be executed repeatedly
		}

	// Loop Control:
		-> loops can be controlled using continue & Break statement

*/

package main

import "fmt"

// Traditional For loop:
func tradition() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// For loop as while loop:
func whileStyle() {
	i := 1

	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
}

// Infinite for loop:
func infiniteFor() {
	i := 1

	for {
		if i == 7 {
			break
		} else {
			fmt.Print(i, " ")
			i++
		}
	}

	fmt.Println()

}

// for...range loop
func forRange() {

	slice := []int{1, 2, 3, 4, 5, 6, 7}

	// Iterate in range:
	for index, val := range slice {
		fmt.Print(index, ":", val, "  ")
	}

	fmt.Println()

	// If we don't want to use index, we can pass _ underscore it will ignore the use.
	for _, v := range slice {
		fmt.Print(v, " ")
	}

	fmt.Println()
}

func main() {
	tradition()
	whileStyle()
	infiniteFor()
	forRange()
}
