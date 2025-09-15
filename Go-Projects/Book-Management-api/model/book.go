/*

// Defining the model Book:
	-> What any books can contains:
		- ID			-> Unique Book ID
		- Title			-> Book Title
		- Author		-> Author name
		- Year			-> releasing year

	-> Here "Book" -> 'B' is in uppercase, which denotes it can be accessed to other packages..
		-> Also their values like "Id", "Title".. First character is in uppercase which denotes the accessibility of values
	-> We can access any package using => import "github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/model"


*/

package model

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
