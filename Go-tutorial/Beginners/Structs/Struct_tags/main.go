/*

// Struct with tags:
	-> It's generally when we are using JSON, DB, etc
	-> Struct tags in go provide a machinist to attach metadata to the fields of a struct.
	-> This metadata is stored as a "string literal" and is typically used by external packages or the reflect package to alter behavior ro provide additional information about the filed.
	-> Using backtick we specify the type & required fields..

*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	u := User{
		Username: "sahilwep",
		Email:    "sahilwep@gmail.com",
	}

	data, _ := json.Marshal(u)
	fmt.Println(string(data))
}
