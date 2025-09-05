/*

// Methods attached to Struct:
	-> Struct have some attributes/property, but we can attach them with some behaviors
	-> During function definition, prior to function name, inside parenthesis (shortHandName StructName)
	-> Example: Syntax

				type Person struct {
					// ... Details of person
				}

				func (p Person) FuncName(args1, args2,...) {
					// logic...
				}

	-> by default struct is passed by value, if we want to refer to original value of struct, pass with pointer & change it
	-> NOTE: struct will automatically defer to the pointer value, we don't need to explicitly denotes.

*/

package main

import (
	"fmt"
	"strconv"
)

// Defining struct:
type Person struct {
	Name string
	Age  int
}

// Method with value receiver: Define a method with receiver
func (p Person) SayHello() string {
	return "hello, my name is " + p.Name + " and i am " + strconv.Itoa(p.Age) + " years old"
}

// Method with pointer receiver: If method need to modify original struct use pointer receiver.
func (p *Person) IncrementAge() {
	p.Age++
}

func main() {

	p1 := Person{
		Name: "Sahil",
		Age:  23,
	}
	fmt.Println(p1)
	fmt.Println(p1.SayHello())
	p1.IncrementAge() // increment age
	fmt.Println(p1)

}
