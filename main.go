package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	keller := author{FirstName: "Gary", LastName: "Keller"}
	// filima := author{FirstName: "John", LastName: "Kilima"}

	books = append(books, book{ID: "1", ISBN: "0001", Title: "The 1 thing", Author: &keller})
	// books = append(books, book{ID: "2", ISBN: "0002", Title: "Here and Go", Author: &filima})

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

type book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *author `json:"author"`
}

type author struct {
	FirstName string `json:"string"`
	LastName  string `json:"string"`
}

var books []book

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
