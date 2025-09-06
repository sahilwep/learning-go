/*

// Recursion:
	-> Function call itself until it not hit any base case...
	-> Let's understand it with very lame printing table example :/


*/

package main

import "fmt"

func PrintTable(i, n int) {
	if i <= 10 {
		fmt.Println(n, "*", i, "=", n*i)
		PrintTable(i+1, n)
	}
}

func main() {
	start, num := 1, 6
	PrintTable(start, num)
}
