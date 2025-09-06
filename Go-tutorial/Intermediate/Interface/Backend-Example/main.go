/*

// Backend Use-case:
	-> Imagine you’re building a database layer.
	-> You want your code to work with MySQL today,
	-> but maybe PostgreSQL tomorrow — without rewriting everything.

	// Explanations:
		-> Like, earlier i say as long as it knows it be done, the contract will done anyhow.

*/

package main

import "fmt"

// Contract (interface):
type Database interface {
	Connect() error
	Query(q string) string
}

// MySql Implementation:
type MySql struct{}

func (MySql) Connect() error {
	fmt.Println("MySql Connected")
	return nil
}

func (MySql) Query(q string) string {
	return "MySql result for query: " + q
}

// Postgresql Implementations:
type PostgreSQL struct{}

func (PostgreSQL) Connect() error {
	fmt.Println("PostgreSQL Connected")
	return nil
}

func (PostgreSQL) Query(q string) string {
	return "PostgreSQL result for Query: " + q
}

// Business logic depends on interface not concrete type
func fetchUser(db Database) {
	db.Connect()
	fmt.Println(db.Query("SELECT * FROM users"))
}

func main() {
	// Create database of interface Database type:
	var db Database

	// using MySQL:
	db = MySql{} // Use MySQL
	fetchUser(db)

	// using PostgreSQL
	db = PostgreSQL{} // Switch Postgresql easily
	fetchUser(db)

}
