package main

// endocing/json is actually run our server when we create and we dont' need to install.
// log: which will log errors.
// net/http package to work with HTTP and that;s we use to create api stuff.
// math/rand to generate id as a random number.
// strconv is convert string.
// go get -u github.com/gorilla/mux is to create a router.
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init books var as a slice book struct.
var books []Book

// Get all books function (handleFunc).
// w http.ResponseWriter it was same when using nodejs app.get("", res).
// r *http.Request it was sama when using nodejs is means req.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get book by id.
func getBooksById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// mux is our router.
	params := mux.Vars(r) // get params.
	// loop through books and find with id.
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create a new book.
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// try using capital letter.
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe.
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update books by id.
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			// try using capital letter.
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// delete books by id.
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init router.
	r := mux.NewRouter()

	// Mock data (todo implement database).
	books = append(books, Book{ID: "1", Isbn: "2343", Title: "Book One", Author: &Author{Firstname: "Julia", Lastname: "Veronica"}})

	// router handlers / endpoints.
	r.HandleFunc("/api/show", getBooks).Methods("GET")
	r.HandleFunc("/api/show/{id}", getBooksById).Methods("GET")
	r.HandleFunc("/api/create", createBooks).Methods("POST")
	r.HandleFunc("/api/update/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/delete/{id}", deleteBooks).Methods("DELETE")
	// make a port pass in the router.
	// use log.Fatal: if it fails we get a throw an error.
	log.Fatal(http.ListenAndServe(":8080", r))
	// if we run right now it will be error because all these route handlers -
	// with functions but we don't have the actual functions yet.
}
