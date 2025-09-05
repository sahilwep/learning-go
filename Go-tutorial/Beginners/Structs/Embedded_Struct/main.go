package main

import "fmt"

type Address struct {
	City    string
	Pin     int64
	Country string
}

type Employee struct {
	Id      int
	Name    string
	Address // embedded
}

func main() {
	e := Employee{
		Id:   1,
		Name: "Sahil",
		Address: Address{
			City:    "Vrindavan",
			Pin:     281121,
			Country: "India",
		},
	}

	fmt.Println(e.Name)
	fmt.Println(e.City)
	fmt.Println(e.Pin)
}
