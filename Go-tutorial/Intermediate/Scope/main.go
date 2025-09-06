/*

// Scope:
	-> It's a block of space where accessibility of any creations valid.
	-> Inside block will by-default access all the outer scope subjects...
	-> But outSide block will not access any inside Scope created subjects..

*/

package main

import "fmt"

// Can be accessed anywhere:
var name string = "sahil"

func main() {
	// Can be accessed inside main scope:
	var num1 int = 1

	if true {
		// can be accessed only inside if block
		var num2 int = 2
		fmt.Println(num2)
		fmt.Println(num1) // accessing num1
		fmt.Println(name) // accessing name
	}
}
