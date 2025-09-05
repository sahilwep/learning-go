/*
// Constant:

	-> Declare with const keyword instead of var
	-> Immutable
	-> Usually declared with uppercase
	-> can be declare inside or outside the function.
*/
package main

import "fmt"

const username string = "sahilwep" // can't be change once it declare

func main() {

	const name string = "Sahil"
	fmt.Println(name, username)

}
