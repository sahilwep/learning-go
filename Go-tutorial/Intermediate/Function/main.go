/*

// Function:
	-> A block of code write to perform intended tasks
	-> we can define by using `func` keyword.

		func FunctionName(args..) returntype {
			// logic...
		}


*/

package main

import (
	"errors"
	"fmt"
)

// Normal greet function
func greet(name string) string {
	return "hello " + name
}

// Add function with multiple parameter
func add(a int, b int) int {
	return a + b
}

// Shorthand: same type
func multiply(a, b, c int) int {
	return a * b * c
}

// Multiple Return value:
func divide(x, y int) (int, error) {
	if y == 0 {
		// We can create an error message using fmt/errors both
		return 0, errors.New("Cannot divide by zero")
		// return 0, fmt.Errorf("Cannot divide by zero")
	}

	return x / y, nil // divide & return nil as error
}

// Named Return value: make sure to use it carefully because a small difference in their name can make whole code error.
func rectangle(w, h int) (area int, perimeter int) {
	area = w * h
	perimeter = 2 * (w * h)
	return // auto return area & perimeter, as we have named them already
}

// Variadic Function: Can accept n number of arguments while function call:
func sum(nums ...int) int { // Inside nums is a slice

	var res int = 0
	for _, val := range nums {
		res += val
	}

	return res
}

// Pass by value:
func update(x int) {
	x += 2
	fmt.Println("by value inside:", x)
}

// Pass By reference: used pointer to access the original address & then modify on address...
func updateNow(x *int) {
	*x += 2
	fmt.Println("by ref inside:", *x)
}

// Function as Value: Function are first-class-citizen -> can be assign, stored, or passed around
func solve(a, b int) int {
	return a + b
}

// Closure: Function can capture variable from their surrounding scope:	here we are using return type as function & their function return type is int
func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {

	// Normal Function
	fmt.Println(greet("sahil"))

	// Function with multiple parameter:
	fmt.Println(add(2, 3))
	fmt.Println(multiply(2, 3, 1))

	// Function with multiple return values:
	res, err := divide(2, 0)
	if err != nil {
		fmt.Println(res)
	} else {
		fmt.Print(res)
	}

	// Function with named return values:
	a, p := rectangle(4, 3)
	fmt.Println("area:", a, "perimeter:", p)

	// Variadic Function:
	println(sum(1, 2))
	println(sum(1, 2, 3))
	println(sum(1, 2, 3, 4))

	// Pass By Value:
	x := 2
	update(x)
	fmt.Println("by value outside", x)

	// Pass By Reference
	updateNow(&x)
	fmt.Println("by ref outside", x)

	// Function as Value:
	op := solve           // solve() is passed into op
	fmt.Println(op(1, 2)) // now we are calling op() value as function

	// Anonymous function (Lambdas):
	double := func(x int) int {
		return x * 2
	}
	fmt.Println(double(5))

	// Inline Anon-func:
	result := func(a, b int) int { return a + b }(2, 3) // here we can calling function at the end of scope using parenthesis
	fmt.Println(result)

	// Closure: Function can capture variable from their surrounding scope => Useful for stateful functions (like generators, caches).
	c := counter() // counter is called once & now c becomes function itself

	fmt.Println(c()) // count = 1
	fmt.Println(c()) // count = 2
	fmt.Println(c()) // count = 3

}
