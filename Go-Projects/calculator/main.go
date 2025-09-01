package main

import (
	"fmt"
)

func input() float32 {
	var temp float32
	fmt.Print("\n[=] Enter the number: ")
	fmt.Scan(&temp)
	return temp
}

func Add(res *float32) {
	x := input()
	*res += x
}

func Sub(res *float32) {
	x := input()
	*res -= x
}

func Mul(res *float32) {
	x := input()
	*res *= x
}

func Div(res *float32) {
	x := input()
	if x != 0 {
		*res /= x
	} else {
		println("\n\tCannot Divide by '0'")
	}
}

func main() {

	var res float32 = 0 // this will be our result

	for {
		// Get input for operations:
		var op int
		fmt.Println("----------------------------")
		fmt.Println("\n\t\t[=] Res: ", res,
			"\n\n[+] Addition: 1",
			"\n[+] Subtract: 2",
			"\n[+] Multiply: 3",
			"\n[+] Division: 4",
			"\n[+] Exit: 5")
		fmt.Println("----------------------------")

		fmt.Print("\n[=] Enter the Operations: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			Add(&res)
		case 2:
			Sub(&res)
		case 3:
			Mul(&res)
		case 4:
			Div(&res)
		case 5:
			return
		default:
			fmt.Print("\n\t{Invalid Input}\n")
		}
	}

}
