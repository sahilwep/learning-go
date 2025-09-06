/*

// Interface:
	-> Interface is like contract in go
	-> It doesn't say how to do something - It's just says what must be done

	// Example:
		A remote control has button like PowerOn(), VolumeUp(), volumeDown().
		A remote doesn't care if it's controlling a TV, AC, or a Projector.
		As long as the device implements those function, the remote can control it, That's what an interface is in Go.

	// Use Case:
		-> Abstraction: Code focused on "what to do" not "how it's done".
		-> Flexibility: different types (struct) can satisfy the same interface.
		-> Testability: You can replace real implementations with mock ones easily.
		-> Plug & Play: Swap implementation without changing the rest of the code.

	// Code Explanations:
		Printer is the contract (Print() method must exist).
		Both Book and Newspaper satisfy the contract.
		Code using Printer doesn’t care whether it’s a Book or Newspaper.

	// Super Simple Analogy:
		Interface = job descriptions (eg: Driver must have Drive() skills").
		Struct = Person (eg: "Sahil", "Prince")
		If Sahil know how to Drive(), he can be hired as a Driver.
		The company (your code) only cares if the person can Drive() - not about his age, background, or anything else.

	// Layman's Terms:
		-> Whenever you carate a variable, we must specify their type
		-> With an interface, we can create a variable of interface type
		-> Now, that variable can hold any struct-methods as long as those struct's methods signature are same with our defined interface methods signature.


	// Backend-Use Case:
		-> Imagine you are building a database layer. you want your code to work with MySQL today, but maybe PostgreSQL tomorrow Without rewriting everything.
		-> Also it's very useful During testing with dummy database, not playing with core database..


*/

package main

import "fmt"

// Define an Interface (Contract)
type Printer interface {
	Print() // NOTE: Signature should be taken carefully, because as long as it is same with struct type, compiler automatically get to know about the interface that we are using
}

// A Concrete type
type Book struct {
	Title string
}

// Book Implements Printer
func (b Book) Print() {
	fmt.Println("Book:", b.Title)
}

// Another Concrete type
type Newspaper struct {
	Name string
}

// Newspaper Implements Printer
func (n Newspaper) Print() {
	fmt.Println("Newspaper:", n.Name)
}

func main() {
	// use Interface Not specific Type: It can hold any struct
	var p Printer

	p = Book{Title: "great king prince"}
	p.Print() // call Book Print()

	p = Newspaper{Name: "A day in a life of Great prince"}
	p.Print() // call Newspaper Print()

}
