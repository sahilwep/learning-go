package main

import "fmt"

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello " + u.Name
}

func main() {
	u1 := User{Name: "Sahil"}
	fmt.Println(u1.Greet())

}
