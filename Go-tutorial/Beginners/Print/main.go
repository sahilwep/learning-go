package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	// we can print anything using Print() or Println() function
	a, b, c := 1, 2, 3
	fmt.Print("Hello, We will learn Printing output\n")
	fmt.Println("This is print value like", a, b, c)

	// We can print anything using Printf(), which is formatted like C Style
	num := 1
	name := "Sahil"
	var ch rune = 65 // rune type used to represent Unicode values
	fmt.Printf("integer: %d, string: %s, char: %c, Type: %T\n", num, name, ch, ch)

	// Sprint(): used to format value and return the resulting string.
	message := fmt.Sprint([]int{1, 2, 3}) // passing slice of integer into sprint() -> this will convert it into string & return
	fmt.Println(message)
	fmt.Printf("message Type: %T\n", message)

	// Log any error value:
	err := errors.New("This is the error message that we are creating") // creating error message
	log.Fatal(err)
}
