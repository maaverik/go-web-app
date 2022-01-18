package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// w.Write([]byte("<h1>Hello world</h1>")) -> this is called internally by Fprintf()
		fmt.Fprintf(w, "<h1>Hello world</h1>")
	} else if r.URL.Path == "/page2" {
		fmt.Fprintf(w, "<h1>Other page</h1>")
	} else {
		http.Error(w, "404: page not found", http.StatusNotFound)
	}
}

// the name main is essential for this function and package to run
func main() {
	// matches all paths
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
