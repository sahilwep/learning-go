/*

// Input:
	-> We can take input using 'scan' or "scanf" formatted function
	-> Scan() & Scanf return value that we have input & error -> This is why we have handel error & define input() & inputf() function.

*/

package main

import (
	"fmt"
	"log"
)

// Normal Input:
func input() {
	var num int
	fmt.Print("Enter a number: ")

	if _, err := fmt.Scan(&num); err != nil {
		log.Fatal(err) // Handel potential error during input
	}

	fmt.Println("Your ID is:", num)
}

// Formatted input:
func inputF() {
	var name string
	fmt.Print("Enter your name: ")

	if _, err := fmt.Scanf("%s", &name); err != nil {
		log.Fatal(err) // Handel potential error during input
	}

	fmt.Printf("Welcome %s\n", name)
}

func main() {
	input()
	inputF()
}
