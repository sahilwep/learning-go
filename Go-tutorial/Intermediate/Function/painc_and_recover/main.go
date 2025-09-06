package main

import "fmt"

// Using Defer we can execute block of recover function if anything panic() happened.
func risky() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r, ", now back to normal")
		}
	}()

	fmt.Println("Everything fine..")
	fmt.Println("Everything fine..")
	fmt.Println("Everything fine..")

	panic("Something went wrong, fatal error19191") // Whenever the panic happen, this message goes into error.

}
func main() {

	risky()
	fmt.Println("Program Continues...")
}
