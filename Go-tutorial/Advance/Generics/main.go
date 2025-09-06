/*
// Generics:
	-> Normally, Go function & types work with specific types(eg, int, string)
	-> Generic let your write a single function/type that works with multiple type without repeating code.
	-> Introduced in go 1.18

	// Syntax:
		func Identity[T any](x T) T {
			return x
		}



	// Without Generic you have to write repeated code
		func sumInts(nums []int) int {...}
		func sumFloat64(nums []float64) float64 {...}

	// With Generics:
		func sum[T int | float64](nums []T) T {
			// do the logic..
		}


*/

package main

import "fmt"

// Sum Example:
func sum[T int | float32](nums []T) T {
	var total T

	for _, val := range nums {
		total += val
	}

	return total
}

// If we are not sure to use which one: We can pass type as "any" similar to interface{}
func PrintAll[T any](nums []T) {
	for _, i := range nums {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// We can Restrict Generic Types which one to use: like options we have {int | float64}
type Number interface {
	int | float64
}

func add[T Number](a, b T) T {
	return a + b
}

// Generics Struct: We can create struct of Generic Type:
type Box[T any] struct {
	Value T
}

func main() {

	// Sum Example:
	fmt.Println(sum([]int{1, 3}))
	fmt.Println(sum([]float32{1.1, 3.2}))

	// Any Type Generics:
	PrintAll([]int{1, 2, 3, 4})
	PrintAll([]float32{1.9, 2.3, 3.34, 4.0})
	PrintAll([]string{"ram", "Shyam", "mohan", "govind"})

	// Restricted type Generic:
	fmt.Println(add(1, 4))
	fmt.Println(add(1.2, 4.4))

	// Generic Struct:
	intBox := Box[int]{Value: 5}
	stringBox := Box[string]{Value: "Items inside box"}
	fmt.Println(intBox)
	fmt.Println(stringBox)

}
