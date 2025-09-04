/*

// Observation:
	-> For testing we created a function to print "hello world"
	-> Then we have to create a file with same name with _test.go name
	-> Inside that we have created a test function which takes only one parameter
	-> then we compare got & want values



*/

package main

import "fmt"

func Hello() string {
	return "hello world"
}

func main() {
	fmt.Println(Hello())
}
