package main

import (
	"fmt"
	"net/http"

	// "\github.com\henrique\cobraStuff\cmd"
	_ "github.com/spf13/cobra"
)

// https://github.com/spf13/cobra:
// You will additionally define flags

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("server starting ...")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
