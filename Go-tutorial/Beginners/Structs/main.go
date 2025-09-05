/*
// Structs:
	-> Structs are collection of fields like record, object or class without methods
	-> Useful to group related data together.

	type structName struct {
		... values
	}

*/

package main

import "fmt"

// Declare struct:
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Zero Values:
	var u1 User
	fmt.Println(u1) // { 0 "" ""}

	// Using Struct literals:
	u2 := User{
		Name:  "Sahil",
		Age:   23,
		Email: "sahlwep@gmail.com",
	}
	fmt.Println(u2)

	// Partial Initialization:
	u3 := User{Name: "Prince"}
	u3.Age = 23
	fmt.Println(u3)

	// Access & update field:
	u2.Name = "राधाकृपाप्राणदास"
	fmt.Println(u2)

	// Pointer to struct: Struct can be use with pointer
	u4 := &User{Name: "Radhkripaprandas"}
	u4.Age = 23 // go will automatically dereference struct pointer when accessing
	fmt.Println(u4.Age)

	// For more about struct -> Visit nested directory
}
