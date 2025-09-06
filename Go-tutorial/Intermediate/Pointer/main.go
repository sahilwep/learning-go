/*
// Pointer:
	-> Pointer is used to store address
	-> We can define pointer type variable and store any address inside that
	-> We can dereference to the original value using '*'
	-> Pointer is very useful specially when we need to perform modification on original value

*/

package main

import "fmt"

// Update function will update original given value:
func Update(age *int, val int) {
	*age += val
}

// Struct Example:
type User struct {
	Name  string
	Age   int
	Email string
}

// using Reference of struct to change original value, without reference we can't change the original value
func (u *User) IncrementAge() {
	u.Age += 1 // we don't need to specify * before that, because struct will automatically handel these things..
}

func main() {
	// Use case of pointer:
	age := 10
	fmt.Println("age before:", age)
	Update(&age, 2)
	fmt.Println("age after:", age)

	// We can also use pointer in struct:
	u1 := User{
		Name:  "Sahil",
		Age:   23,
		Email: "sahilwep@gmail.com",
	}

	fmt.Println(u1)
	u1.IncrementAge()
	fmt.Println(u1)

}
