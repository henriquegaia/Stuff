package main

import (
	"fmt"
	"net/http"

	arithmetic "github.com/henrique/cobraStuff/arithmetic"
)

func main() {
	fmt.Println(arithmetic.Ints(1,2,3))
	// http.HandleFunc("/", index)
	// http.HandleFunc("/about", about)
	// fmt.Println("server starting ...")
	// http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
