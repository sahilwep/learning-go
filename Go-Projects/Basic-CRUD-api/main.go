/*

// Creating CRUD api using Go-Gorilla-mux

	// Create:
		-> Create function will create random id into movie slice

	// Read:
		-> Read Function will fetch movie with requested ID

	// Update:
		-> Because we are not using Database:
			-> we will first delete the movie that user send the ID
			-> then simply add a new movie struct that user send.
			-> This is not the right way to work when we are doing this into database.

	// Delete:
		-> Delete function will delete that requested ID


// Testing:

-> Go and play with http methods & perform crud applications...

`JSON
{
    "id": "",
    "isbn": "23743",
    "title": "Movie One",
    "director": {
        "f_name": "John",
        "l_name": "Wick"
    }
}
`


*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create struct of movie details:
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Create struct of directors details
type Director struct {
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
}

var movies []Movie // create a slice movie of struct type

// All CRUD Functions:
func getMovies(w http.ResponseWriter, r *http.Request) { // w = response that we want to send, r = request pointer
	w.Header().Set("Content-Type", "application/json") // Set header
	json.NewEncoder(w).Encode(movies)                  // encode movie into json & then send
}

func deleteMovie(w http.ResponseWriter, r *http.Request) { // Delete movie function:
	w.Header().Set("Content-Type", "application/json") // set header
	params := mux.Vars(r)                              // get the params from the request 'r'

	// Iterate on movie list:
	for index, items := range movies {
		if items.ID == params["id"] { // match if movie id is equal with passed id:
			movies = append(movies[:index], movies[index+1:]...) // Using append() we are deleting items from slice -> Basically it will get the data from that index & in place of that it will append all the items next to that slice..
			break                                                // and done
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) { // get a particular movie from the slice -> encode it into json & send it
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterate & find it from the slice:
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) { // function used to create movie:
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)    // User send some data into json -> Convert it into format readable by golang
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // create a random number b/w range 1 to that given number & then typecast into string & assign to movie.Id
	movies = append(movies, movie)                // append that incoming requested movie into movies slice
	json.NewEncoder(w).Encode(movie)              // return user that movie has been created
}

func updateMovie(w http.ResponseWriter, r *http.Request) { // update movie function

	w.Header().Set("Content-Type", "application/json") // Set json content type:
	params := mux.Vars(r)                              // Access the params

	// Loop over the movie, range
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // delete the movie with the id that user send

			// add a new movie - the movie that user send in the body
			var movie Movie                            // create a new movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // pass the reference into the movie by decoding it from json
			movie.ID = params["id"]                    // we will make sure the ID should be same as previous user send id
			movies = append(movies, movie)             // now append new incoming movie into our movies slice
			json.NewEncoder(w).Encode(movie)           // return user that your movie is updated
			return
		}
	}
}

func main() {
	r := mux.NewRouter() // this will be our router

	// Initialize our server with some movies:
	movies = append(movies, Movie{ID: "1", Isbn: "23743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Wick"}})
	movies = append(movies, Movie{ID: "2", Isbn: "47271", Title: "Movie Two", Director: &Director{FirstName: "Peter", LastName: "Phillip"}})

	// Routes:
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server:
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
