package main

import (
	"fmt"
)

func input() float32 {
	var temp float32
	for {
		fmt.Print("\n[=] Enter the number: ")
		if _, err := fmt.Scan(&temp); err != nil {
			fmt.Println("Invalid Input")
			continue
		}
		return temp
	}

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
func Power(res *float32) {

	var x int
	fmt.Print("\n[=] Enter the number: ")
	fmt.Scan(&x)

	var temp float32 = 1
	for i := 1; i <= x; i++ {
		temp = temp * *res
	}

	*res = temp
}

func main() {

	var res float32 = 0 // this will be our result
	res = input()

	for {
		// Get input for operations:
		var op int
		fmt.Println("----------------------------")
		fmt.Println("\n\t\t[=] Res: ", res,
			"\n\n[+] Addition: 1",
			"\n[+] Subtract: 2",
			"\n[+] Multiply: 3",
			"\n[+] Division: 4",
			"\n[+] Power: 5",
			"\n[+] Exit: 6")
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
			Power(&res)
		case 6:
			fmt.Println("Thanks for Using Calculator")
			return
		default:
			fmt.Print("\n\t{Invalid Input}\n")
		}
	}

}
