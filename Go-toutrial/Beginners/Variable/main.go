/*

// Variable:
	-> variable name should start with any letter or _ underscore
	-> Can't start with any digit
	-> Can't have space b/w them
	-> Case sensitive
	-> can't use go keywords for name convention

*/

package main

import "fmt"

var name = "Sahil" // this way we can also declare variable

func main() {

	fmt.Println(name)

	// var keyword following by variable name & their type
	var fruit string = "apple"
	fmt.Println(fruit)

	// Use := sign used by variable name to declare & assign value to it.
	college := "LPU"
	idCard := 1
	fmt.Println(college, idCard)

	// Assign Type + type inferred
	var car string = "BMW" // type is string
	var driver = "Sahil"   // type is inferred
	race := true           // type is inferred
	fmt.Println(driver, "Drives", car, "for race", race)

	// Go multiple Variable declarations:
	var one, two, three, four int = 1, 2, 3, 4
	println(one, two, three, four)

	// Variable Declaration in block:
	var (
		id        int    = 0
		number           = 1
		firstName string = "Sahil"
		lastName  string = "Sharma"
	)
	fmt.Println(id, number, firstName, lastName)

}
