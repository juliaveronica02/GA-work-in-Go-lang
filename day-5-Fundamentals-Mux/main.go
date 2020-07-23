package main

// endocing/json is actually run our server when we create and we dont' need to install.
// log: which will log errors.
// net/http package to work with HTTP and that;s we use to create api stuff.
// math/rand to generate id as a random number.
// strconv is convert string.
// go get -u github.com/gorilla/mux is to create a router.
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct (model).
// a struct should have an iron is kind of like class it's used for object oriented programming -
// in golang.
// to define struct we use type.
type Book struct {
	// set up properties.
	// id in lowerCase this will be the json field that we fetch.
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// author have own struct (asterisk author).
	Author *Author `json:"author"`
}

// seperate struct for the author.
// author  struct (we can add another else).
type Author struct {
	firstName string `json:"firstname"`
	lastName  string `json:"firstname"`
}

// Get all books function (handleFunc).
func getBooks() {

}

func main() {
	// init router.
	r := mux.NewRouter()

	// router handlers / endpoints.
	r.HandleFunc("api/books", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBooksById).Methods("GET")
	r.HandleFunc("api/books", createBooks).Methods("POST")
	r.HandleFunc("api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBooks).Methods("DELETE")
	// make a port pass in the router.
	// use log.Fatal: if it fails we get a throw an error.
	log.Fatal(http.ListenAndServe(":8080", r))
	// if we run right now it will be error because all these route handlers -
	// with functions but we don't have the actual functions yet.
}
