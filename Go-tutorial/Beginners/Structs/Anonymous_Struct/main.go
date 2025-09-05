/*

// Anonymous Struct:
	-> Anonymous struct is struct type defined without name
	-> This allows for the creations of temporary, one-off data structure without need to declare a formal type.
	-> Anonymous structs are particularly useful in scenarios where a specific struct is only relevant within a limited scope, promoting cleaner and more concise code.


*/

package main

import "fmt"

func main() {
	// Declaring and initializing anonymous struct:
	Sahil := struct {
		FirstName string
		LastName  string
		Age       int
		email     string
	}{
		FirstName: "Sahil",
		LastName:  "Sharma",
		Age:       23,
		email:     "sahilwep@gmail.com",
	}

	fmt.Println(Sahil.FirstName)
	fmt.Println(Sahil.Age)

	// Another Example: Embedding an anonymous struct
	data := struct {
		ID   int
		info struct { // anonymous struct as field
			Description string
			value       float32
		}
	}{
		ID: 1,
		info: struct {
			Description string
			value       float32
		}{
			Description: "This is First Product ",
			value:       1830.57,
		},
	}

	fmt.Println(data.ID)
	fmt.Println(data.info.Description)
	fmt.Println(data.info.value)
}
