/*

// Enumeration:
	-> An Enum (Enumerations type) is a way to give names to names to related constant value.
	-> It helps you to represent a set of fixed options(states, nodes, categories) in your code.
	-> It doesn't have keywords enum, but we simulates it using constants + iota

	// Without enums, you might write something like:
		-> Example:
			status := 1	// what does 1 means??

		-> Without enum it's hard to read, instead enums let you do:

	// With Enums:
		const (
			Pending = iota
			Approved
			Reject
		)

		Status := Approved

*/

package main

import "fmt"

// Create Enums for OrderStatus
type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
)

func main() {
	Status := Shipped

	if Status == Shipped {
		fmt.Println("Your Order is Shipped")
	}

	// If we increment status value => it will go to next value:
	Status++
	if Status == Delivered {
		fmt.Println("Your Order is Delivered")
	}

}
