/*
// Numbers:

	-> In go we can use numbers with using:

		int  int8  int16  int32  int64				-> Sign integer that can handel given bytes
		uint uint8 uint16 uint32 uint64 uintptr		-> Unsigned integer, handel given bytes

		byte 										-> alias for uint8

		rune 										-> alias for int32, represents a Unicode code point

		float32 float64								-> Handel floating values

		complex64 complex128						-> Store complex numbers like iota
*/
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// Integer:
	var num1 int = 183
	var num2 int32 = 182828
	fmt.Printf("num1 has %T type, and num2 has %T type\n", num1, num2)

	// Byte Type:
	var num3 byte = 65 // a = 65 in ASCII
	fmt.Printf("Num3 is of %T type & value in character is: %c\n", num3, num3)

	// Rune => ASCII is limited, we can store extra symbol like '世', so that it's used & it's simple an alias for int32 used to represent unicode
	var r rune = '世' // represent unicode character
	fmt.Printf("rune: %c, has value: %d, and it's type is: %T\n", r, r, r)

	// Floating numbers:
	var decVal float64 = 18383.1812232
	var decVal2 float32 = 1822.223131
	fmt.Println(decVal, "&", decVal2)

	// Complex Type:
	var result complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(result)
}
