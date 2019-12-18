package main

import (
	"fmt"
	"log"
	"net/http"

	// "str/conv"
	// "math/rand"
	// "encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))

	// http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
	// fmt.Println("server starting ...")
	// http.ListenAndServe(":3000", nil)
}

type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
