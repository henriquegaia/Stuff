package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "str/conv"
	// "math/rand"
	// "encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "2", ISBN: "45345", Title: "The 1 thing",
		Author: &Author{FirstName: "Gary", LastName: "Keller"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("server starting ...")
	http.ListenAndServe(":3000", r)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
}

type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"string"`
	LastName  string `json:"string"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request)    {}
func createBook(w http.ResponseWriter, r *http.Request) {}
func updateBook(w http.ResponseWriter, r *http.Request) {}
func deleteBook(w http.ResponseWriter, r *http.Request) {}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
