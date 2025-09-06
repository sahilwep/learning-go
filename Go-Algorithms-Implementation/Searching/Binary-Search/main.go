package main

import (
	"fmt"
	"sort"
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

// Binary Search: TC: O(log(n))
func BinarySearch(nums []int, n, target int) int {
	low, high := 0, n-1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // if not found
}

func main() {
	// take Input:
	num, n, target := input()

	sort.Ints(num) // sort th given input
	fmt.Println("Array after sorting:", num)

	if i := BinarySearch(num, n, target); i != -1 {
		fmt.Println("Found at index:", i)
	} else {
		fmt.Println("Not Found")
	}
}
