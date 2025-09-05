package main

import "fmt"

// Arithmetic Operators: Return numeric result based on given operation
func arithmetic() {
	fmt.Println("Arithmetic: ")
	fmt.Println(2 + 2)  // addition +
	fmt.Println(4 - 2)  // subtraction -
	fmt.Println(2 * 2)  // multiplication *
	fmt.Println(4 / 2)  // Division /
	fmt.Println(16 % 2) // modulus % => gives remainder

	// Increment & decrement:
	a := 1
	a++ // increment by 1
	a-- // decrement by 1
	fmt.Println(a)
	fmt.Println()
}

// Relational Operators: Return -> true / false
func relational() {
	fmt.Println("Relational: ")
	fmt.Println(2 > 2)  // grater than
	fmt.Println(2 >= 2) // greater or equal
	fmt.Println(2 < 2)  // less than
	fmt.Println(2 <= 2) // less or equal
	fmt.Println(2 == 2) // equal
	fmt.Println(2 != 2) // not-equal
	fmt.Println()
}

// Logical Operators: Compare conditions on left & right & return true/false
func logical() {
	fmt.Println("Logical: ")
	fmt.Println(4 > 3 && 4 < 10) // compare both side & if both return true => then only return true : else false
	fmt.Println(4 > 5 || 4 < 10) // compare both side & if any return true => return true
	fmt.Println(!(4 == 4))       // invert the given result
	fmt.Println()

}

// Bitwise Operators: works at bit level or used to perform bit by bit operations
func bitwise() {
	fmt.Println("Bitwise: ")
	fmt.Println(2 & 4)  // takes two number & does AND on bit of two number => return the result afterwards
	fmt.Println(2 | 4)  // takes two number & does OR on bit of two number => return the result afterwards
	fmt.Println(2 ^ 3)  // takes two number & does XOR on bit of two number	=> return the result afterwards
	fmt.Println(2 << 1) // takes two number & left shift the bit of first number, second number decide the number of bit place shift => Here 4 is left shift by 1 bit position
	fmt.Println(4 >> 1) // takes two number & right shift the bit of first number, second number decide the number of bit place shift

	// Bit Clear: &^
	/*
		In this example,
		if num is 10110101 and positionToClear is 2, the mask would be 00000100.
		The &^ operation would then clear the bit at that position in num.

			10110101	-> Num
			00000100	-> Mask	=> This we can create by creating a int variable = 1 & left shift 2 position
		Resulting in 10110001.
	*/

	var num byte = 0b10110101    // example number
	positionToClear := 2         // clear bit position at 2 (0-based indexing)
	mask := 1 << positionToClear // Create a mask with a 1 at the desired position and 0s elsewhere => mask: 00000100
	result := num &^ byte(mask)

	fmt.Printf("Original number: %08b, Changed to %08b\n", num, result)
	fmt.Println()
}

// Assignment Operators
func assignment() {
	fmt.Println("Assignment: ")
	var a, b int = 10, 20 // assignment	 '='

	// Assignment works on Arithmetic:
	a += 1
	b -= 1
	a *= 2
	a /= 2
	b %= a

	// Assignment works on Bitwise:
	a &= 1
	b ^= 2
	a |= 4
	a <<= 2
	b >>= 1
	fmt.Println(a, b)
	fmt.Println()
}

// Misc Operators:
func misc() {
	fmt.Println("Operators:  &   *   <-   ")

	// Address Resolution operator: &
	var num int = 10
	fmt.Println("value:", num, "and address is:", &num)

	// Pointer: *
	var addr *int = &num
	fmt.Println("pointer store address of var", *addr, "value inside pointer is:", addr)

	// <- This operator is known as Receive Operators
	ch := make(chan int) // create a channel of integer
	go func() {          // create a goroutine using "go"keyword & define function
		ch <- 7 // send a value to the channel in separate goroutine.
	}() // last call the function with ()

	value := <-ch // receive value from channel
	fmt.Println("Received value", value, "from channel")

}

func main() {

	arithmetic()
	relational()
	logical()
	bitwise()
	assignment()
	misc()

}
