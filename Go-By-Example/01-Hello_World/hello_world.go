/*

	- For any program we are writing, we have to import "package main" & inside "func main(){}"
	- For printing anything on console, we can use "fmt.Println()" function which comes from "fmt" package.


	- To run program directly we can use: `go run fileName.go`
	- To build our program we can use: `go build filename.go`
	- To execute the binary: `./fileName`

*/

package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
