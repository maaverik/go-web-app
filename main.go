package main

import (
	"fmt"
	"go-web-app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404: page not found</h1>")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// the name main is essential for this function and package to run
func main() {
	usersController := controllers.NewUsers()
	staticController := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticController.HomeView).Methods("GET")
	r.Handle("/contact", staticController.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)

	http.ListenAndServe(":3000", r)
}
