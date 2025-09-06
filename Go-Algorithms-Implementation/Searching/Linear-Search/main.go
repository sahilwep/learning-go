package main

import (
	"fmt"
)

// Input Function:
func input() ([]int, int, int) {
	// Take size of array & target
	var n, target int
	fmt.Scan(&n, &target)

	// Take input for slice:
	num := make([]int, n) // create a slice of length 'n'
	for i, _ := range num {
		fmt.Scan(&num[i])
	}

	return num, n, target
}

// Linear Search: TC: O(n)
func LinearSearch(nums []int, n, target int) int {

	for index, val := range nums {
		if val == target {
			return index
		}
	}

	return -1 // if not found
}

func main() {
	// take Input:
	num, n, target := input()

	if i := LinearSearch(num, n, target); i != -1 {
		fmt.Println("Found at index:", i)
	} else {
		fmt.Println("Not Found")
	}
}
