/*

// Conditional:
	-> if else are same as C/C++, it's just we don't have to specify the parenthesis () to specify the conditions.
	-> Note: At the end of if '}' else should be start, if we needed, In next line we can't start else...
		-> It's small convention from go




*/

package main

import "fmt"

func main() {

	a, b, c := 30, 20, 40

	if a > b {
		fmt.Println("a is grater than b,")
		if a > c {
			fmt.Println("a is also grater than c")
		} else {
			fmt.Println("but a is not grater than c")
		}
	} else if b > a {
		fmt.Println("b is grater than a")
	} else {
		fmt.Println("both are equal")
	}

	num1 := 20
	if res := num1 - 10; res != num1 {
		fmt.Println("We can also use declaration & Checking inside 'if'")
	}
}
