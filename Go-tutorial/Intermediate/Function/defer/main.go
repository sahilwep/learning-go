/*
// defer in function:
	-> "defer" is used to schedule a function call to be executed immediately before the surrounding function returns.
	-> This means the defer function will run wether the surrounding function return normally, due to an error or because of panic.
	-> How defer works:
		-> Delayed Execution: the primary purus is to delay the execution of a function until the very end of enclosing function's scope.
		-> Resource management: it is commonly used for resource cleanup, such as closing files, releasing locks, or closing network connections..
		-> LIFO Order: it follows the LIFO order -> last defer will executed first
		-> Argument Evaluations: the arguments to a defer function are evaluated at a time the defer statement is execute, not when the deferred function is actually called. This is an important distinction to remember when dealing with variable that might change before the defer function runs.
		-> Error handling and panics: defer particularly useful in conjunction with the panic and recover for robust error handling


*/

package main

import (
	"fmt"
	"os"
)

// Example of File Handling using defer:
func fileWriteExample() {

	file, err := os.Create("Create.txt")
	if err != nil {
		fmt.Println("Error Creating files: ", err)
		return
	}
	defer file.Close() // this will ensure file is closed when main exit

	_, err = file.WriteString("hello, we are Go Defer learning")
	if err != nil {
		fmt.Println("Error while writing file")
		return
	}

	fmt.Println("File Written successfully")
}

func main() {
	defer fmt.Println("Line 1")
	fmt.Println("Line 2")

	fileWriteExample()
}
