/*

// Slice:
	-> Slice is dynamic, flexible view into an underlying array.
	-> Unlike array, slice can grow or shrink
	-> Internally, a slice is describe by three things:
		-> Pointer: points to underlying array
		-> length: number of elements in the slice
		-> Capacity: maximum number of elements(before reallocation)

	-> Slice are referenced, so modifying them may affect original array.
	-> append(), copy() & slice = num[:]
	-> We have two thing length & capacity -> length is the value inside the array slice present & capacity is the total capacity, it capture in memory & if the capacity is full, then it will double, like vector in C++


*/

package main

import "fmt"

func main() {
	// Empty Slice:
	var s1 []int
	fmt.Println(s1)        // []
	fmt.Println(s1 == nil) // true (nil slice)

	// Slice literal:
	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)

	// Declare Using make(length & capacity)
	s3 := make([]int, 3, 5) // length = 3, capacity = 5
	fmt.Println(s3)         // [0 0 0]

	// Length & capacity:
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(len(nums)) // 5
	fmt.Println(cap(nums)) // 5

	// now when we slice:
	sub := nums[1:4] // elements index 1 to 3
	fmt.Println(sub)
	fmt.Println(len(sub)) // 3
	fmt.Println(cap(sub)) // 5 (from index 1 to end of nums)

	// Append Elements:
	s := []int{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Println(s)

	// Append Slice:
	sl1 := []int{1, 2, 3}
	sl2 := []int{4, 5, 6}
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)

	// Copy Slice:
	a := []int{1, 2, 3}
	b := make([]int, len(a)) // make a empty slice of length of 'a'
	copy(b, a)               // copy a into b
	fmt.Println(b)

	// Iterating slice:
	names := []string{"Radha", "Krishna", "Radharaman", "Radhavallav", "Bihari ji", "Gopinath"}
	for _, name := range names {
		fmt.Print(name, " ")
	}
	fmt.Println()

	// Modify Slice:
	values := []int{1, 2, 3, 4, 5}
	values[3] = 0
	fmt.Println(values)

	// slice from slice:
	flower := []string{"sunflower", "rose", "lily", "lavender"}
	favFlower := flower[1:3]
	fmt.Println(favFlower)

	// Reslice:
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[:3])
	fmt.Println(x[2:])
	fmt.Println(x[:])

	// Nil vs empty slice:
	var p1 []int  // nil slice
	p2 := []int{} // empty slice
	fmt.Println(p1 == nil)
	fmt.Println(p2 == nil)

}
