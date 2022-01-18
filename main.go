package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact page</h1>")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404: page not found</h1>")
}

// the name main is essential for this function and package to run
func main() {
	// matches all paths
	r := mux.NewRouter()
	h := http.HandlerFunc(pageNotFound)
	r.NotFoundHandler = h
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
