/*
// Switch Case:

	-> Switch case don't required any break statement-
	-> It's very similar to any other language like C/C++
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	var day int

	// Read input:
	fmt.Print("Enter the number: ")
	if _, err := fmt.Scan(&day); err != nil {
		log.Fatal(err) // Handel potential error during input
	}

	// Switch-Case:
	switch day {
	case 1:
		fmt.Println("day 1, Monday")
	case 2:
		fmt.Println("day 2, Tuesday")
	case 3:
		fmt.Println("day 3, Wednesday")
	case 4:
		fmt.Println("day 4, Thursday")
	case 5:
		fmt.Println("day 5, Friday")
	case 6:
		fmt.Println("day 6, Saturday")
	case 7:
		fmt.Println("day 7, Sunday")
	default:
		fmt.Println("Invalid input")
	}
}
